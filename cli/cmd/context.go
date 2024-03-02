package cmd

import (
	"github.com/spf13/cobra"
	"go.datalift.io/admiral/common/client"
)

type ContextCmd struct {
	Cmd *cobra.Command
}

func NewContextCmd(clientOpts *client.Options) *ContextCmd {
	root := &ContextCmd{}
	var delete bool

	cmd := &cobra.Command{
		Use:     "context",
		Short:   "Switch between contexts",
		Aliases: []string{"ctx"},
		RunE: func(cmd *cobra.Command, args []string) error {
			//localCfg, err := localconfig.ReadLocalConfig(clientOpts.ConfigPath)
			return nil
		},
	}

	cmd.Flags().BoolVar(&delete, "delete", false, "Delete the context instead of switching to it")

	root.Cmd = cmd
	return root
}
