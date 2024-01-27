package main

import (
	"go.datalift.io/datalift/server/cmd/assets"
	"go.datalift.io/datalift/server/gateway"
	"go.datalift.io/datalift/server/service"
)

var MockServiceFactory = service.Factory{}

func main() {
	cf := gateway.CoreComponentFactory

	// Replace core services with any available mocks.
	cf.Services = MockServiceFactory

	gateway.Run(gateway.ParseFlags(), cf, assets.VirtualFS)
}
