package server

import (
	"log"

	"github.com/spf13/cobra"
)

type MigrateCmd struct {
	Cmd *cobra.Command
}

func NewMigrateCmd() *MigrateCmd {
	root := &MigrateCmd{}

	cmd := &cobra.Command{
		Use:           "migrate",
		Short:         "Database migration tool",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			log.Fatal("Not implemented")
		},
	}

	root.Cmd = cmd
	return root
}
