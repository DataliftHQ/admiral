package cmd

import (
	"errors"
	"fmt"
	"go.datalift.io/admiral/cli/cmd/account"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	cobracompletefig "github.com/withfig/autocomplete-tools/integrations/cobra"
	utillog "go.datalift.io/admiral/cli/internal/util/log"

	"go.datalift.io/admiral/cli/internal/common"
	"go.datalift.io/admiral/common/client"
	"go.datalift.io/admiral/common/config"
	"go.datalift.io/admiral/common/util/env"
	"go.datalift.io/admiral/common/util/text"
	"go.datalift.io/admiral/common/version"
)

func Execute(version version.Version, exit func(int), args []string) {
	newRootCmd(version, exit).Execute(args)
}

type MultiError struct {
	Errors []error
}

// Error implements the error interface for MultiError.
func (me *MultiError) Error() string {
	var errorStrings []string
	for _, err := range me.Errors {
		errorStrings = append(errorStrings, err.Error())
	}
	return strings.Join(errorStrings, ", ")
}

func (cmd *rootCmd) Execute(args []string) {
	cmd.cmd.SetArgs(args)

	if err := cmd.cmd.Execute(); err != nil {
		code := 1
		msg := "command failed"
		eerr := &exitError{}
		if errors.As(err, &eerr) {
			code = eerr.code
			if eerr.details != "" {
				msg = eerr.details
			}
		}
		log.WithError(err).Error(msg)
		cmd.exit(code)
	}
}

type rootCmd struct {
	cmd  *cobra.Command
	exit func(int)
}

func newRootCmd(version version.Version, exit func(int)) *rootCmd {
	var logFormat, logLevel string
	var clientOpts client.Options

	root := &rootCmd{
		exit: exit,
	}

	cmd := &cobra.Command{
		Use:           "admiral",
		Short:         "Admiral - Platform Orchestrator",
		Version:       version.String(),
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := SetLogFormat(logFormat); err != nil {
				return wrapError(err, "failed to set log format")
			}

			if err := SetLogLevel(logLevel); err != nil {
				return wrapError(err, "failed to set log level")
			}

			return nil
		},
	}
	cmd.SetVersionTemplate("{{.Version}}")

	// general options
	cmd.PersistentFlags().BoolP("help", "h", false, "help for admiral cli")
	cmd.PersistentFlags().StringVar(&logFormat, "logformat", "text", "Set the logging format. One of: text|json")
	cmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "Set the logging level. One of: debug|info|warn|error")

	// config options
	defaultConfigFile, err := config.DefaultPath()
	if err != nil {
		log.WithError(err).Fatal("failed to get default config file")
	}
	cmd.PersistentFlags().StringVarP(&clientOpts.ConfigFile, "config", "f", defaultConfigFile, "Path to configuration file")

	// auth options
	cmd.PersistentFlags().StringVar(&clientOpts.AccessToken, "access-token", "", "access token")

	// server options
	cmd.PersistentFlags().StringVarP(&clientOpts.ServerAddress, "server", "s", env.StringFromEnv(common.EnvServerAddress, ""), "host:port of the api server")
	cmd.PersistentFlags().BoolVar(&clientOpts.PlainText, "plaintext", false, "disable tls")
	cmd.PersistentFlags().BoolVarP(&clientOpts.Insecure, "insecure", "i", false, "skip server certificate and domain verification")
	cmd.PersistentFlags().StringVar(&clientOpts.CertFile, "server-crt", "", "server certificate file")

	// client options
	cmd.PersistentFlags().StringVar(&clientOpts.ClientCertFile, "client-crt", "", "client certificate file")
	cmd.PersistentFlags().StringVar(&clientOpts.ClientCertKeyFile, "client-crt-key", "", "client certificate key file")
	cmd.PersistentFlags().IntVar(&clientOpts.HttpRetryMax, "http-retry-max", 0, "maximum number of retries to establish http connection to server")
	cmd.PersistentFlags().StringSliceVarP(&clientOpts.Headers, "header", "H", []string{}, "Sets additional header to all requests. (Can be repeated multiple times to add multiple headers, also supports comma separated headers)")

	cmd.AddCommand(
		account.NewAccountCmd(&clientOpts).Cmd,
		NewLoginCmd(&clientOpts).Cmd,
		NewLogoutCmd(&clientOpts).Cmd,
		NewFooCmd(&clientOpts).Cmd,
		newManCmd().cmd,
		cobracompletefig.CreateCompletionSpecCommand(),
	)

	root.cmd = cmd
	return root
}

func SetLogFormat(logFormat string) (err error) {
	switch strings.ToLower(logFormat) {
	case utillog.JsonFormat:
		err = os.Setenv(common.EnvLogFormat, utillog.JsonFormat)
	case utillog.TextFormat, "":
		err = os.Setenv(common.EnvLogFormat, utillog.TextFormat)
	default:
		err = fmt.Errorf("unknown log format '%s'", logFormat)
	}
	log.SetFormatter(utillog.CreateFormatter(logFormat))
	return err
}

func SetLogLevel(logLevel string) (err error) {
	level, err := log.ParseLevel(text.FirstNonEmpty(logLevel, log.InfoLevel.String()))
	if err != nil {
		return err
	}
	err = os.Setenv(common.EnvLogLevel, level.String())
	log.SetLevel(level)
	return err
}
