package common

import "github.com/coreos/go-oidc/v3/oidc"

const (
	EnvLogFormat     = "ADMIRAL_LOG_FORMAT"
	EnvLogLevel      = "ADMIRAL_LOG_LEVEL"
	EnvServerAddress = "ADMIRAL_SERVER_ADDRESS"
)

var (
	DefaultOAuth2Scopes = []string{oidc.ScopeOpenID, oidc.ScopeOfflineAccess}
)
