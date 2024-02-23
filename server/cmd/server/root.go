package server

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"go.datalift.io/admiral/common/version"
)

func Execute(version version.Version, exit func(int), args []string) {
	newRootCmd(version, exit).Execute(args)
}

func newRootCmd(version version.Version, exit func(int)) *rootCmd {
	root := &rootCmd{
		exit: exit,
	}

	cmd := &cobra.Command{
		Use:               "admiral-server",
		Short:             "Admiral Platform Orchestrator Server",
		Version:           version.String(),
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}
	cmd.SetVersionTemplate("{{.Version}}")

	cmd.AddCommand(
		NewMigrateCmd().Cmd,
		NewRunCmd().Cmd,
	)

	root.cmd = cmd
	return root
}

type rootCmd struct {
	cmd  *cobra.Command
	exit func(int)
}

func (cmd *rootCmd) Execute(args []string) {
	cmd.cmd.SetArgs(args)

	if shouldPrependRun(cmd.cmd, args) {
		cmd.cmd.SetArgs(append([]string{"run"}, args...))
	}

	if err := cmd.cmd.Execute(); err != nil {
		log.WithError(err).Error("command failed")
		cmd.exit(1)
	}
}

func shouldPrependRun(cmd *cobra.Command, args []string) bool {
	xmd, _, _ := cmd.Find(args)
	if xmd != cmd {
		return false
	}

	if len(args) != 1 {
		return true
	}

	for _, s := range []string{"-h", "--help", "-v", "--version"} {
		if s == args[0] {
			return false
		}
	}

	return true
}
