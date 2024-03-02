package account

import (
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	sessionv1 "go.datalift.io/admiral/common/api/session/v1"
	"go.datalift.io/admiral/common/client"
	"go.datalift.io/admiral/common/util/io"
)

type userInfoCmd struct {
	cmd  *cobra.Command
	opts userInfoOpts
}

type userInfoOpts struct {
}

func newGetUserInfoCmd(clientOpts *client.Options) *userInfoCmd {
	root := &userInfoCmd{}
	var (
		output string
	)
	cmd := &cobra.Command{
		Use:               "info",
		Short:             "Get user info",
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, sessionClient := client.NewClientOrDie(clientOpts).NewSessionClientOrDie()
			defer io.Close(conn)

			response, err := sessionClient.GetSession(cmd.Context(), &sessionv1.GetSessionRequest{})
			if err != nil {
				return err
			}

			switch output {
			case "yaml":
				yamlBytes, err := yaml.Marshal(response)
				if err != nil {
					return err
				}

				fmt.Println(string(yamlBytes))
			case "json":
				jsonBytes, err := json.MarshalIndent(response, "", "  ")
				if err != nil {
					return err
				}

				fmt.Println(string(jsonBytes))
			case "":
				fmt.Printf("Logged In: %v\n", response.LoggedIn)
				if response.LoggedIn {
					fmt.Printf("Username: %s\n", response.Username)
					fmt.Printf("Issuer: %s\n", response.Iss)
					fmt.Printf("Groups: %v\n", strings.Join(response.Groups, ","))
				}
			default:
				log.Fatalf("Unknown output format: %s", output)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "Output format. One of: yaml, json")
	root.cmd = cmd
	return root
}
