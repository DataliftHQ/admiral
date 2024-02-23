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
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/hashicorp/go-retryablehttp"
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
	UserAgent         string
	PlainText         bool
	Insecure          bool
	CertFile          string
	ClientCertFile    string
	ClientCertKeyFile string
	Headers           []string
	HttpRetryMax      int

	AccessToken    string
	RefreshToken   string
	CredentialPath string

	Issuer   string
	ClientId string
	Scopes   []string
}

type Client interface {
	ClientOptions() Options
	HTTPClient() (*http.Client, error)
	OIDCConfig(context.Context) (*oauth2.Config, *oidc.Provider, error)

	NewSessionClient() (io.Closer, sessionv1.SessionAPIClient, error)
	NewSessionClientOrDie() (io.Closer, sessionv1.SessionAPIClient)
	NewSettingsClient() (io.Closer, settingsv1.SettingsAPIClient, error)
	NewSettingsClientOrDie() (io.Closer, settingsv1.SettingsAPIClient)
}

type client struct {
	Options

	httpClient  *http.Client
	CertPEMData []byte
	ClientCert  *tls.Certificate
}

func (c *client) ClientOptions() Options {
	return Options{
		ServerAddress: c.ServerAddress,
		PlainText:     c.PlainText,
		Insecure:      c.Insecure,
	}
}

// fix
func NewClient(opts *Options) (Client, error) {
	var c client

	if opts.AccessToken == "" {

	}

	// get credential file

	if opts.Issuer != "" {
		c.Issuer = opts.Issuer
	}
	if c.Issuer == "" {
		return nil, errors.New("openid provider url is unspecified")
	}

	if opts.ClientId != "" {
		c.ClientId = opts.ClientId
	}
	if c.ClientId == "" {
		return nil, errors.New("oauth client id is unspecified")
	}

	if opts.ServerAddress != "" {
		c.ServerAddress = opts.ServerAddress
	}
	if c.ServerAddress == "" {
		return nil, errors.New("server address is unspecified")
	}

	if opts.CertFile != "" {
		b, err := os.ReadFile(opts.CertFile)
		if err != nil {
			return nil, err
		}
		c.CertPEMData = b
	}

	if opts.ClientCertFile != "" && opts.ClientCertKeyFile != "" {
		clientCert, err := tls.LoadX509KeyPair(opts.ClientCertFile, opts.ClientCertKeyFile)
		if err != nil {
			return nil, err
		}
		c.ClientCert = &clientCert
	} else if opts.ClientCertFile != "" || opts.ClientCertKeyFile != "" {
		return nil, errors.New("--client-crt and --client-crt-key must always be specified together")
	}

	if opts.PlainText {
		c.PlainText = true
	}
	if opts.Insecure {
		c.Insecure = true
	}

	if opts.HttpRetryMax > 0 {
		retryClient := retryablehttp.NewClient()
		retryClient.RetryMax = opts.HttpRetryMax
		c.httpClient = retryClient.StandardClient()
	} else {
		c.httpClient = &http.Client{}
	}

	if !c.PlainText {
		tlsConfig, err := c.tlsConfig()
		if err != nil {
			return nil, err
		}
		c.httpClient.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	if opts.Scopes == nil || len(opts.Scopes) == 0 {
		c.Scopes = []string{oidc.ScopeOpenID}
	} else {
		c.Scopes = opts.Scopes
	}

	if opts.UserAgent == "" {
		c.UserAgent = fmt.Sprintf("%s/%s", "admiral-client", "unknown")
	} else {
		c.UserAgent = opts.UserAgent
	}

	err := c.refreshAccessToken()
	if err != nil {
		return nil, err
	}

	c.Headers = opts.Headers

	return &c, nil
}

func NewClientOrDie(opts *Options) Client {
	client, err := NewClient(opts)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// fix
func (c *client) refreshAccessToken() error {

	return nil
}

// fix
func (c *client) redeemRefreshToken(t *oauth2.Token) (*oauth2.Token, error) {
	httpClient, err := c.HTTPClient()
	if err != nil {
		return nil, err
	}

	ctx := oidc.ClientContext(context.Background(), httpClient)

	oauth2conf, _, err := c.OIDCConfig(ctx)
	if err != nil {
		return nil, err
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

	headers, err := parseHeaders(c.Headers)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		headers.Set("User-Agent", c.UserAgent)
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

// fix
func (c *client) OIDCConfig(ctx context.Context, set settingsv1.SettingsResponse) (*oauth2.Config, *oidc.Provider, error) {
	var clientID string
	var issuerURL string
	var scopes []string

	if set.OidcConfig != nil && set.OidcConfig.Issuer != "" {
		if set.OidcConfig.CliClientId != "" {
			clientID = set.OidcConfig.CliClientId
		} else {
			clientID = set.OidcConfig.ClientId
		}
		issuerURL = set.OidcConfig.Issuer
		scopes = set.OidcConfig.Scopes
	} else {
		return nil, nil, fmt.Errorf("%s is not configured with SSO", c.ServerAddress)
	}

	provider, err := oidc.NewProvider(ctx, issuerURL)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to query provider %q: %v", issuerURL, err)
	}

	oidcConf, err := oidcutil.ParseConfig(provider)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to parse provider config: %v", err)
	}

	scopes = oidcutil.GetScopesOrDefault(scopes)
	if oidcutil.OfflineAccess(oidcConf.ScopesSupported) {
		scopes = append(scopes, oidc.ScopeOfflineAccess)
	}

	oauth2conf := oauth2.Config{
		ClientID: clientID,
		Scopes:   scopes,
		Endpoint: provider.Endpoint(),
	}

	return &oauth2conf, provider, nil
}

func (c *client) newConn() (*grpc.ClientConn, io.Closer, error) {
	closers := make([]io.Closer, 0)
	serverAddr := c.ServerAddress
	network := "tcp"

	var creds credentials.TransportCredentials
	if !c.PlainText {
		tlsConfig, err := c.tlsConfig()
		if err != nil {
			return nil, nil, err
		}
		creds = credentials.NewTLS(tlsConfig)
	}

	endpointCredentials := jwtCredentials{
		Token: c.AccessToken,
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

	headers, err := parseHeaders(c.Headers)
	if err != nil {
		return nil, nil, err
	}
	for k, vs := range headers {
		for _, v := range vs {
			ctx = metadata.AppendToOutgoingContext(ctx, k, v)
		}
	}

	if c.UserAgent != "" {
		dialOpts = append(dialOpts, grpc.WithUserAgent(c.UserAgent))
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
		log.Fatalf("Failed to establish connection to %s: %v", c.ServerAddress, err)
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
		log.Fatalf("Failed to establish connection to %s: %v", c.ServerAddress, err)
	}
	return conn, setIf
}
