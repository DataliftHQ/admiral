package main

import (
	"go.datalift.io/admiral/server/cmd/assets"
	"go.datalift.io/admiral/server/gateway"
	"go.datalift.io/admiral/server/service"
)

var MockServiceFactory = service.Factory{}

func main() {
	cf := gateway.CoreComponentFactory

	// Replace core services with any available mocks.
	cf.Services = MockServiceFactory

	gateway.Run(gateway.ParseFlags(), cf, assets.VirtualFS)
}
