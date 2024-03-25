package foo

import (
	"context"
	"errors"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"

	foov1 "go.datalift.io/admiral/common/api/foo/v1"
	"go.datalift.io/admiral/server/endpoint"
	"go.datalift.io/admiral/server/service"
	"go.datalift.io/admiral/server/service/temporal"
)

const (
	Name = "admiral.endpoint.foo"
)

func New(_ *anypb.Any, log *zap.Logger, scope tally.Scope) (endpoint.Endpoint, error) {
	t, ok := service.Registry["admiral.service.temporal"]
	if !ok {
		return nil, errors.New("could not find service")
	}

	mgr, ok := t.(temporal.ClientManager)
	if !ok {
		return nil, errors.New("service was not the correct type")
	}

	client, err := mgr.GetNamespaceClient("default")
	if err != nil {
		return nil, err
	}

	return &endp{
		temporalClient: client,
		logger:         log,
		scope:          scope,
	}, nil
}

type endp struct {
	temporalClient temporal.Client

	logger *zap.Logger
	scope  tally.Scope
}

func (e *endp) Register(r endpoint.Registrar) error {
	foov1.RegisterFooAPIServer(r.GRPCServer(), e)
	return r.RegisterJSONGateway(foov1.RegisterFooAPIHandler)
}

func (a *endp) GetFoo(ctx context.Context, _ *foov1.GetFooRequest) (*foov1.GetFooResponse, error) {
	return &foov1.GetFooResponse{
		Foo: "Bar",
	}, nil
}
