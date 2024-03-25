package settings

import (
	"context"
	"github.com/uber-go/tally/v4"
	settingsv1 "go.datalift.io/admiral/common/api/settings/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"

	"go.datalift.io/admiral/server/endpoint"
)

const (
	Name = "admiral.endpoint.settings"
)

func New(_ *anypb.Any, log *zap.Logger, scope tally.Scope) (endpoint.Endpoint, error) {
	return &endp{
		settings: newSettingsAPI(),
		logger:   log,
		scope:    scope,
	}, nil
}

type endp struct {
	settings settingsv1.SettingsAPIServer
	logger   *zap.Logger
	scope    tally.Scope
}

func (e *endp) Register(r endpoint.Registrar) error {
	settingsv1.RegisterSettingsAPIServer(r.GRPCServer(), e.settings)
	return r.RegisterJSONGateway(settingsv1.RegisterSettingsAPIHandler)
}

func newSettingsAPI() settingsv1.SettingsAPIServer {
	return &settingsAPI{}
}

type settingsAPI struct{}

func (a *settingsAPI) Settings(context.Context, *settingsv1.SettingsRequest) (*settingsv1.SettingsResponse, error) {
	return &settingsv1.SettingsResponse{
		Url: "http://localhost:8080",
		OidcConfig: &settingsv1.OIDCConfig{
			Name:        "admiral",
			Issuer:      "http://localhost:9090/realms/admiral",
			ClientId:    "admiral",
			CliClientId: "admiral-cli",
			Scopes:      []string{"openid", "offline_access", "profile", "email"},
		},
	}, nil
}
