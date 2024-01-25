package client

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

	httputil "go.datalift.io/datalift/common/util/http"
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

	//NewUserClient() (io.Closer, usergrpc.UserAPIClient, error)
	//NewUserClientOrDie() (io.Closer, usergrpc.UserAPIClient)
	//NewSettingsClient() (io.Closer, settingsv1.SettingsAPIClient, error)
	//NewSettingsClientOrDie() (io.Closer, settingsv1.SettingsAPIClient)
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
		c.UserAgent = fmt.Sprintf("%s/%s", "datalift-client", "unknown")
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

func (c *client) refreshAccessToken() error {

	return nil
}

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

func NewClientOrDie(opts *Options) Client {
	client, err := NewClient(opts)
	if err != nil {
		log.Fatal(err)
	}
	return client
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

func (c *client) OIDCConfig(ctx context.Context) (*oauth2.Config, *oidc.Provider, error) {
	provider, err := oidc.NewProvider(ctx, c.Issuer)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query provider %q: %v", c.Issuer, err)
	}

	endpoint := provider.Endpoint()
	endpoint.AuthStyle = oauth2.AuthStyleInParams

	oauth2conf := oauth2.Config{
		ClientID: c.ClientId,
		Endpoint: endpoint,
		Scopes:   c.Scopes,
	}

	return &oauth2conf, provider, nil
}
