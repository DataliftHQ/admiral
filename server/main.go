package main

import (
	"go.datalift.io/datalift/server/cmd/assets"
	"go.datalift.io/datalift/server/gateway"
)

func main() {
	flags := gateway.ParseFlags()
	components := gateway.CoreComponentFactory

	gateway.Run(flags, components, assets.VirtualFS)
}
