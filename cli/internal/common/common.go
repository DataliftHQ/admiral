package common

import "github.com/coreos/go-oidc/v3/oidc"

const (
	EnvLogFormat     = "ADMIRAL_LOG_FORMAT"
	EnvLogLevel      = "ADMIRAL_LOG_LEVEL"
	EnvServerAddress = "ADMIRAL_SERVER_ADDRESS"

	EnvOAuth2Issuer   = "DATALIFT_OAUTH2_ISSUER"
	EnvOAuth2ClientId = "DATALIFT_OAUTH2_CLIENT_ID"
	EnvOAuth2Scopes   = "DATALIFT_OAUTH2_SCOPES"

	DefaultServerAddress  = "api.datalift.io:443"
	DefaultOAuth2Issuer   = "https://auth.datalift.io"
	DefaultOAuth2ClientId = "44972cb5-9739-4b8d-ac29-6dcccca3e9db"
)

var (
	DefaultOAuth2Scopes = []string{oidc.ScopeOpenID, oidc.ScopeOfflineAccess}
)
