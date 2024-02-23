package account

import (
	"github.com/spf13/cobra"

	"go.datalift.io/admiral/common/client"
)

type AccountCmd struct {
	Cmd  *cobra.Command
	opts accountOpts
}

type accountOpts struct {
}

func NewAccountCmd(clientOpts *client.Options) *AccountCmd {
	root := &AccountCmd{}

	cmd := &cobra.Command{
		Use:           "account",
		Short:         "Manage account settings",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
	}

	cmd.AddCommand(
		newGetUserInfoCmd(clientOpts).cmd,
	)

	root.Cmd = cmd
	return root
}
