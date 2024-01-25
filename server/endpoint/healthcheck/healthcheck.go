package healthcheck

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"

	healthcheckv1 "go.datalift.io/datalift/common/api/healthcheck/v1"
	"go.datalift.io/datalift/server/endpoint"
)

const (
	Name = "datalift.module.healthcheck"
)

func New(*any.Any, *zap.Logger, tally.Scope) (endpoint.Endpoint, error) {
	endp := &endp{
		api: newAPI(),
	}
	return endp, nil
}

type endp struct {
	api healthcheckv1.HealthcheckAPIServer
}

func (e *endp) Register(r endpoint.Registrar) error {
	healthcheckv1.RegisterHealthcheckAPIServer(r.GRPCServer(), e.api)
	return r.RegisterJSONGateway(healthcheckv1.RegisterHealthcheckAPIHandler)
}

func newAPI() healthcheckv1.HealthcheckAPIServer {
	return &healthcheckAPI{}
}

type healthcheckAPI struct{}

func (a *healthcheckAPI) Healthcheck(context.Context, *healthcheckv1.HealthcheckRequest) (*healthcheckv1.HealthcheckResponse, error) {
	return &healthcheckv1.HealthcheckResponse{}, nil
}
