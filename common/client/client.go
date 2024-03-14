package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"go.datalift.io/admiral/common/config"
	"io"
	"math"
	"net"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	sessionv1 "go.datalift.io/admiral/common/api/session/v1"
	settingsv1 "go.datalift.io/admiral/common/api/settings/v1"
	"go.datalift.io/admiral/common/util/env"
	grpcutil "go.datalift.io/admiral/common/util/grpc"
	httputil "go.datalift.io/admiral/common/util/http"
	ioutil "go.datalift.io/admiral/common/util/io"
	oidcutil "go.datalift.io/admiral/common/util/oidc"
)

const (
	MetaDataTokenKey        = "token"
	EnvAdmiralServer        = "ADMIRAL_SERVER"
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
}

func NewClient(opts *Options) (Client, error) {
	var c client
	var err error

	c.config, err = config.Read(opts.ConfigFile)
	if err != nil {
		return nil, err
	}

	//// config exists, use it and update with options
	//if cfg != nil {
	//	c.ServerAddress = cfg.ServerAddress
	//	if cfg.CACertificateAuthorityData != "" {
	//		c.CertPEMData, err = base64.StdEncoding.DecodeString(cfg.CACertificateAuthorityData)
	//		if err != nil {
	//			return nil, err
	//		}
	//	}
	//
	//	if cfg.ClientCertificateData != "" && cfg.ClientCertificateKeyData != "" {
	//		clientCertData, err := base64.StdEncoding.DecodeString(cfg.ClientCertificateData)
	//		if err != nil {
	//			return nil, err
	//		}
	//		clientCertKeyData, err := base64.StdEncoding.DecodeString(cfg.ClientCertificateKeyData)
	//		if err != nil {
	//			return nil, err
	//		}
	//		clientCert, err := tls.X509KeyPair(clientCertData, clientCertKeyData)
	//		if err != nil {
	//			return nil, err
	//		}
	//		c.ClientCert = &clientCert
	//	} else if cfg.ClientCertificateData != "" || cfg.ClientCertificateKeyData != "" {
	//		return nil, errors.New("ClientCertificateData and ClientCertificateKeyData must always be specified together")
	//	}
	//	c.PlainText = cfg.PlainText
	//	c.Insecure = cfg.Insecure
	//	c.AccessToken = cfg.Token.AccessToken
	//	c.RefreshToken = cfg.Token.RefreshToken
	//}

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

	//// Override auth-token if specified in env variable or CLI flag
	//c.AccessToken = env.StringFromEnv(EnvAdmiralAccessToken, c.AccessToken)
	//if opts.AccessToken != "" {
	//	c.AccessToken = strings.TrimSpace(opts.AccessToken)
	//}
	//
	//if opts.CertFile != "" {
	//	b, err := os.ReadFile(opts.CertFile)
	//	if err != nil {
	//		return nil, err
	//	}
	//	c.CertPEMData = b
	//}
	//
	//if opts.ClientCertFile != "" && opts.ClientCertKeyFile != "" {
	//	clientCert, err := tls.LoadX509KeyPair(opts.ClientCertFile, opts.ClientCertKeyFile)
	//	if err != nil {
	//		return nil, err
	//	}
	//	c.ClientCert = &clientCert
	//} else if opts.ClientCertFile != "" || opts.ClientCertKeyFile != "" {
	//	return nil, errors.New("--client-crt and --client-crt-key must always be specified together")
	//}

	if opts.PlainText {
		c.config.Settings.PlainText = true
	}
	if opts.Insecure {
		c.config.Settings.Insecure = true
	}

	// TODO: blue tape
	if opts.HttpRetryMax > 0 {
		retryClient := retryablehttp.NewClient()
		retryClient.RetryMax = opts.HttpRetryMax
		c.httpClient = retryClient.StandardClient()
	} else {
		c.httpClient = &http.Client{}
	}

	// TODO: blue tape
	if !c.config.Settings.PlainText {
		tlsConfig, err := c.tlsConfig()
		if err != nil {
			return nil, err
		}
		c.httpClient.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	//if cfg != nil {
	//	err = c.refreshAccessToken(cfg, opts.ConfigFile)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//
	//c.Headers = opts.Headers

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

func (c *client) NewSessionClient() (io.Closer, sessionv1.SessionAPIClient, error) {
	conn, closer, err := c.newConn()
	if err != nil {
		return nil, nil, err
	}
	sessionIf := sessionv1.NewSessionAPIClient(conn)
	return closer, sessionIf, nil
}

func (c *client) NewSessionClientOrDie() (io.Closer, sessionv1.SessionAPIClient) {
	conn, sessionIf, err := c.NewSessionClient()
	if err != nil {
		log.Fatalf("Failed to establish connection to %s: %v", c.config.Settings.ServerAddress, err)
	}
	return conn, sessionIf
}

func (c *client) NewSettingsClient() (io.Closer, settingsv1.SettingsAPIClient, error) {
	conn, closer, err := c.newConn()
	if err != nil {
		return nil, nil, err
	}
	setIf := settingsv1.NewSettingsAPIClient(conn)
	return closer, setIf, nil
}

func (c *client) NewSettingsClientOrDie() (io.Closer, settingsv1.SettingsAPIClient) {
	conn, setIf, err := c.NewSettingsClient()
	if err != nil {
		log.Fatalf("Failed to establish connection to %s: %v", c.config.Settings.ServerAddress, err)
	}
	return conn, setIf
}
