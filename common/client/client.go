package client

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	foov1 "go.datalift.io/admiral/common/api/foo/v1"
	sessionv1 "go.datalift.io/admiral/common/api/session/v1"
	settingsv1 "go.datalift.io/admiral/common/api/settings/v1"
	"go.datalift.io/admiral/common/config"
	"go.datalift.io/admiral/common/util/env"
	grpcutil "go.datalift.io/admiral/common/util/grpc"
	httputil "go.datalift.io/admiral/common/util/http"
	ioutil "go.datalift.io/admiral/common/util/io"
	oidcutil "go.datalift.io/admiral/common/util/oidc"
)

const (
	EnvAdmiralAccessToken   = "ADMIRAL_ACCESS_TOKEN"
	EnvAdmiralgRPCMaxSizeMB = "ADMIRAL_GRPC_MAX_SIZE_MB"
)

var (
	MaxGRPCMessageSize = env.ParseNumFromEnv(EnvAdmiralgRPCMaxSizeMB, 200, 0, math.MaxInt32) * 1024 * 1024
)

type Options struct {
	ServerAddress     string
	PlainText         bool
	Insecure          bool
	CertFile          string
	ClientCertFile    string
	ClientCertKeyFile string
	AccessToken       string
	ConfigFile        string
	UserAgent         string
	Headers           []string
	HttpRetryMax      int
}

type client struct {
	config     *config.Config
	httpClient *http.Client
}

type Client interface {
	HTTPClient() (*http.Client, error)

	OIDCConfig(context.Context) (*oauth2.Config, *oidc.Provider, error)
	Config() *config.Config

	NewSessionClient() (io.Closer, sessionv1.SessionAPIClient, error)
	NewSessionClientOrDie() (io.Closer, sessionv1.SessionAPIClient)
	NewSettingsClient() (io.Closer, settingsv1.SettingsAPIClient, error)
	NewSettingsClientOrDie() (io.Closer, settingsv1.SettingsAPIClient)

	NewFooClient() (io.Closer, foov1.FooAPIClient, error)
	NewFooClientOrDie() (io.Closer, foov1.FooAPIClient)
}

// NOTE: Client options are saved automatically, a process that might not suit
// all cases. It might be more user-friendly to allow users to explicitly
// save their settings, or to introduce a configuration option enabling this
// feature. Additionally, if this approach is adopted, it's advisable to store
// credentials (access and refresh tokens) separately from the settings.
func NewClient(opts *Options) (Client, error) {
	var c client
	var err error

	c.config, err = config.Read(opts.ConfigFile)
	if err != nil {
		return nil, err
	}

	if opts.UserAgent == "" {
		c.config.Settings.UserAgent = fmt.Sprintf("%s/%s", "admiral-client", "unknown")
	} else {
		c.config.Settings.UserAgent = opts.UserAgent
	}

	if opts.ServerAddress != "" {
		c.config.Settings.ServerAddress = opts.ServerAddress
	}
	if c.config.Settings.ServerAddress == "" {
		return nil, errors.New("server address is unspecified")
	}

	// Override access-token if specified in env variable or CLI flag
	c.config.Token.AccessToken = env.StringFromEnv(EnvAdmiralAccessToken, c.config.Token.AccessToken)
	if opts.AccessToken != "" {
		c.config.Token.AccessToken = strings.TrimSpace(opts.AccessToken)
	}

	if opts.CertFile != "" {
		b, err := os.ReadFile(opts.CertFile)
		if err != nil {
			return nil, err
		}

		c.config.Settings.CertPEMData = b
	}

	if opts.ClientCertFile != "" && opts.ClientCertKeyFile != "" {
		clientCert, err := tls.LoadX509KeyPair(opts.ClientCertFile, opts.ClientCertKeyFile)
		if err != nil {
			return nil, err
		}
		c.config.Settings.ClientCert = &clientCert
	} else if opts.ClientCertFile != "" || opts.ClientCertKeyFile != "" {
		return nil, errors.New("--client-crt and --client-crt-key must always be specified together")
	}

	if opts.PlainText {
		c.config.Settings.PlainText = true
	}
	if opts.Insecure {
		c.config.Settings.Insecure = true
	}

	if opts.HttpRetryMax > 0 {
		retryClient := retryablehttp.NewClient()
		retryClient.RetryMax = opts.HttpRetryMax
		c.httpClient = retryClient.StandardClient()
	} else {
		c.httpClient = &http.Client{}
	}

	if !c.config.Settings.PlainText {
		tlsConfig, err := c.tlsConfig()
		if err != nil {
			return nil, err
		}
		c.httpClient.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	err = c.refreshAccessToken()
	if err != nil {
		return nil, err
	}

	c.config.Settings.Headers = opts.Headers

	// Save the config file
	err = c.config.Save()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func NewClientOrDie(opts *Options) Client {
	client, err := NewClient(opts)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (c *client) refreshAccessToken() error {
	if c.config.Token.RefreshToken == "" {
		// If we have no refresh token, there's no point in doing anything
		return nil
	}

	parser := jwt.NewParser(jwt.WithoutClaimsValidation())
	var claims jwt.RegisteredClaims
	_, _, err := parser.ParseUnverified(c.config.Token.AccessToken, &claims)
	if err != nil {
		return err
	}

	validator := jwt.NewValidator()
	if validator.Validate(claims) == nil {
		// token is still valid
		return nil
	}

	log.Debug("Auth token no longer valid. Refreshing")
	token, err := c.redeemRefreshToken()
	if err != nil {
		return err
	}

	c.config.Token = *token
	err = c.config.Save()
	if err != nil {
		return err
	}

	return nil
}

func (c *client) redeemRefreshToken() (*oauth2.Token, error) {
	httpClient, err := c.HTTPClient()
	if err != nil {
		return nil, err
	}
	ctx := oidc.ClientContext(context.Background(), httpClient)

	oauth2conf, _, err := c.OIDCConfig(ctx)
	if err != nil {
		return nil, err
	}

	t := &oauth2.Token{
		RefreshToken: c.config.Token.RefreshToken,
	}
	token, err := oauth2conf.TokenSource(ctx, t).Token()
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (c *client) HTTPClient() (*http.Client, error) {
	tlsConfig, err := c.tlsConfig()
	if err != nil {
		return nil, err
	}

	headers, err := parseHeaders(c.config.Settings.Headers)
	if err != nil {
		return nil, err
	}

	if c.config.Settings.UserAgent != "" {
		headers.Set("User-Agent", c.config.Settings.UserAgent)
	}

	return &http.Client{
		Transport: &httputil.TransportWithHeader{
			RoundTripper: &http.Transport{
				TLSClientConfig: tlsConfig,
				Proxy:           http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
			Header: headers,
		},
	}, nil
}

func (c *client) OIDCConfig(ctx context.Context) (*oauth2.Config, *oidc.Provider, error) {
	var clientId string
	var issuerUrl string
	var scopes []string

	settingsConn, settingsClient, err := c.NewSettingsClient()
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = settingsConn.Close() }()

	settings, err := settingsClient.Settings(ctx, &settingsv1.SettingsRequest{})
	if err != nil {
		return nil, nil, err
	}

	if settings.OidcConfig != nil && settings.OidcConfig.Issuer != "" {
		if settings.OidcConfig.CliClientId != "" {
			clientId = settings.OidcConfig.CliClientId
		} else {
			clientId = settings.OidcConfig.ClientId
		}
		issuerUrl = settings.OidcConfig.Issuer
		scopes = settings.OidcConfig.Scopes
	} else {
		return nil, nil, fmt.Errorf("%s is not configured with SSO", c.config.Settings.ServerAddress)
	}

	provider, err := oidc.NewProvider(ctx, issuerUrl)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query provider %q: %v", issuerUrl, err)
	}

	oidcConf, err := oidcutil.ParseConfig(provider)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse provider config: %v", err)
	}

	scopes = oidcutil.GetScopesOrDefault(scopes)
	if oidcutil.OfflineAccess(oidcConf.ScopesSupported) {
		scopes = append(scopes, oidc.ScopeOfflineAccess)
	}

	oauth2conf := oauth2.Config{
		ClientID: clientId,
		Scopes:   scopes,
		Endpoint: provider.Endpoint(),
	}

	return &oauth2conf, provider, nil
}

func (c *client) Config() *config.Config {
	return c.config
}

func (c *client) newConn() (*grpc.ClientConn, io.Closer, error) {
	closers := make([]io.Closer, 0)
	serverAddr := c.config.Settings.ServerAddress
	network := "tcp"

	var creds credentials.TransportCredentials
	if !c.config.Settings.PlainText {
		tlsConfig, err := c.tlsConfig()
		if err != nil {
			return nil, nil, err
		}
		creds = credentials.NewTLS(tlsConfig)
	}

	endpointCredentials := jwtCredentials{
		Token: c.config.Token.AccessToken,
	}

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithMax(3),
		grpcretry.WithBackoff(grpcretry.BackoffLinear(1000 * time.Millisecond)),
	}

	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(endpointCredentials))
	dialOpts = append(dialOpts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxGRPCMessageSize), grpc.MaxCallSendMsgSize(MaxGRPCMessageSize)))
	dialOpts = append(dialOpts, grpc.WithStreamInterceptor(grpcretry.StreamClientInterceptor(retryOpts...)))
	dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(retryOpts...)))
	dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(grpcutil.OTELUnaryClientInterceptor()))
	dialOpts = append(dialOpts, grpc.WithStreamInterceptor(grpcutil.OTELStreamClientInterceptor()))

	ctx := context.Background()

	headers, err := parseHeaders(c.config.Settings.Headers)
	if err != nil {
		return nil, nil, err
	}
	for k, vs := range headers {
		for _, v := range vs {
			ctx = metadata.AppendToOutgoingContext(ctx, k, v)
		}
	}

	if c.config.Settings.UserAgent != "" {
		dialOpts = append(dialOpts, grpc.WithUserAgent(c.config.Settings.UserAgent))
	}

	conn, e := grpcutil.BlockingDial(ctx, network, serverAddr, creds, dialOpts...)
	closers = append(closers, conn)

	return conn, ioutil.NewCloser(func() error {
		var firstErr error
		for i := range closers {
			err := closers[i].Close()
			if err != nil {
				firstErr = err
			}
		}
		return firstErr
	}), e
}

func (c *client) NewSettingsClient() (io.Closer, settingsv1.SettingsAPIClient, error) {
	conn, closer, err := c.newConn()
	if err != nil {
		return nil, nil, err
	}
	settingClient := settingsv1.NewSettingsAPIClient(conn)
	return closer, settingClient, nil
}

func (c *client) NewSettingsClientOrDie() (io.Closer, settingsv1.SettingsAPIClient) {
	conn, settingClient, err := c.NewSettingsClient()
	if err != nil {
		log.Fatalf("Failed to establish connection to %s: %v", c.config.Settings.ServerAddress, err)
	}
	return conn, settingClient
}

// ----------------------------------------------------------------------------

func (c *client) NewSessionClient() (io.Closer, sessionv1.SessionAPIClient, error) {
	conn, closer, err := c.newConn()
	if err != nil {
		return nil, nil, err
	}
	sessionClient := sessionv1.NewSessionAPIClient(conn)
	return closer, sessionClient, nil
}

func (c *client) NewSessionClientOrDie() (io.Closer, sessionv1.SessionAPIClient) {
	conn, sessionClient, err := c.NewSessionClient()
	if err != nil {
		log.Fatalf("Failed to establish connection to %s: %v", c.config.Settings.ServerAddress, err)
	}
	return conn, sessionClient
}

func (c *client) NewFooClient() (io.Closer, foov1.FooAPIClient, error) {
	conn, closer, err := c.newConn()
	if err != nil {
		return nil, nil, err
	}
	fooClient := foov1.NewFooAPIClient(conn)
	return closer, fooClient, nil
}

func (c *client) NewFooClientOrDie() (io.Closer, foov1.FooAPIClient) {
	conn, fooClient, err := c.NewFooClient()
	if err != nil {
		log.Fatalf("Failed to establish connection to %s: %v", c.config.Settings.ServerAddress, err)
	}
	return conn, fooClient
}
