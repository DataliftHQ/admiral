package healthcheck

import (
	"context"

	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"

	healthcheckv1 "go.datalift.io/admiral/common/api/healthcheck/v1"
	"go.datalift.io/admiral/server/endpoint"
)

const (
	Name = "admiral.endpoint.healthcheck"
)

type endp struct {
	api healthcheckv1.HealthcheckAPIServer
}

func New(*anypb.Any, *zap.Logger, tally.Scope) (endpoint.Endpoint, error) {
	endp := &endp{
		api: newHealthcheckAPI(),
	}
	return endp, nil
}

func (e *endp) Register(r endpoint.Registrar) error {
	healthcheckv1.RegisterHealthcheckAPIServer(r.GRPCServer(), e.api)
	return r.RegisterJSONGateway(healthcheckv1.RegisterHealthcheckAPIHandler)
}

func newHealthcheckAPI() healthcheckv1.HealthcheckAPIServer {
	return &healthcheckAPI{}
}

type healthcheckAPI struct{}

func (a *healthcheckAPI) Healthcheck(context.Context, *healthcheckv1.HealthcheckRequest) (*healthcheckv1.HealthcheckResponse, error) {
	return &healthcheckv1.HealthcheckResponse{}, nil
}
