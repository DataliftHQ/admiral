package main

import (
	vers "go.datalift.io/admiral/common/version"
	"go.datalift.io/admiral/server/cmd/server"
	"os"
)

var (
	version = ""
	commit  = ""
	date    = ""
	builtBy = ""
)

func main() {
	server.Execute(
		buildVersion(version, commit, date, builtBy),
		os.Exit,
		os.Args[1:],
	)
}

func buildVersion(version, commit, date, builtBy string) vers.Version {
	return vers.GetVersion(
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
