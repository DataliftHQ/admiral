package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.datalift.io/admiral/common/client"
)

type LogoutCmd struct {
	Cmd  *cobra.Command
	opts logoutOpts
}

type logoutOpts struct {
}

func NewLogoutCmd(clientOpts *client.Options) *LogoutCmd {
	root := &LogoutCmd{}

	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Log out from Admiral",
		RunE: func(cmd *cobra.Command, args []string) error {
			admiral, err := client.NewClient(clientOpts)
			if err != nil {
				return err
			}
			//admiral.Logout()
			fmt.Sprintf("%+v", admiral)
			
			return nil
		},
	}

	root.Cmd = cmd
	return root
}
