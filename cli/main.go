package main

import (
	_ "embed"
	"os"

	"go.datalift.io/datalift/cli/cmd"
	vers "go.datalift.io/datalift/common/version"
)

var (
	version = ""
	commit  = ""
	date    = ""
	builtBy = ""
)

//go:embed art.txt
var asciiArt string

func main() {
	cmd.Execute(
		buildVersion(version, commit, date, builtBy),
		os.Exit,
		os.Args[1:],
	)
}

func buildVersion(version, commit, date, builtBy string) vers.Version {
	return vers.GetVersion(
		vers.WithAsciiArt(asciiArt),
		func(v *vers.Version) {
			if commit != "" {
				v.GitCommit = commit
			}
			if date != "" {
				v.BuildDate = date
			}
			if version != "" {
				v.Version = version
			}
			if builtBy != "" {
				v.BuiltBy = builtBy
			}
		},
	)
}
