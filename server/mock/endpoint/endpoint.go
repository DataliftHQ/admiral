package endpointmock

import (
	"google.golang.org/grpc"

	"go.datalift.io/datalift/server/endpoint"
)

type MockRegistrar struct {
	Server *grpc.Server
}

func (m *MockRegistrar) GRPCServer() *grpc.Server { return m.Server }

func (m *MockRegistrar) RegisterJSONGateway(handlerFunc endpoint.GatewayRegisterAPIHandlerFunc) error {
	return nil
}
