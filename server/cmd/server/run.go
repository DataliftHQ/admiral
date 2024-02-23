package server

import (
	"github.com/spf13/cobra"

	"go.datalift.io/admiral/server/cmd/assets"
	"go.datalift.io/admiral/server/gateway"
)

type RunCmd struct {
	Cmd *cobra.Command
}

func NewRunCmd() *RunCmd {
	root := &RunCmd{}

	cmd := &cobra.Command{
		Use:           "run",
		Short:         "Run the server",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			flags := gateway.ParseFlags()
			components := gateway.CoreComponentFactory

			gateway.Run(flags, components, assets.VirtualFS)
		},
	}

	root.Cmd = cmd
	return root
}
