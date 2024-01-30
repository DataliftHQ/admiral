package main

import (
	"go.datalift.io/admiral/server/cmd/assets"
	"go.datalift.io/admiral/server/gateway"
)

func main() {
	flags := gateway.ParseFlags()
	components := gateway.CoreComponentFactory

	gateway.Run(flags, components, assets.VirtualFS)
}
