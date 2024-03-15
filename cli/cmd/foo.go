package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.datalift.io/admiral/common/client"
	"go.datalift.io/admiral/common/util/io"

	foov1 "go.datalift.io/admiral/common/api/foo/v1"
)

type FooCmd struct {
	Cmd *cobra.Command
}

func NewFooCmd(clientOpts *client.Options) *FooCmd {
	root := &FooCmd{}

	cmd := &cobra.Command{
		Use:   "foo",
		Short: "Execute foo endpoint",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, fooClient := client.NewClientOrDie(clientOpts).NewFooClientOrDie()
			defer io.Close(conn)

			response, err := fooClient.GetFoo(cmd.Context(), &foov1.GetFooRequest{})
			if err != nil {
				return err
			}

			log.Infof("%+v\n", response)
			return nil
		},
	}

	root.Cmd = cmd
	return root
}
