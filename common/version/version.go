package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"text/tabwriter"
	"time"
)

const unknown = "unknown"

type Version struct {
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
	BuildDate string `json:"buildDate"`
	BuiltBy   string `json:"builtBy"`
	GoVersion string `json:"goVersion"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
	AsciiArt  string `json:"-"`
}

type Option func(i *Version)

func WithVersion(version string) Option {
	return func(v *Version) {
		v.Version = version
	}
}

func WithGitCommit(commit string) Option {
	return func(v *Version) {
		v.GitCommit = commit
	}
}

func WithBuiltBy(name string) Option {
	return func(v *Version) {
		v.BuiltBy = name
	}
}

func WithAsciiArt(art string) Option {
	return func(v *Version) {
		v.AsciiArt = art
	}
}

func getBuildInfo() *debug.BuildInfo {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return nil
	}
	return bi
}

func getGitVersion(bi *debug.BuildInfo) string {
	if bi == nil {
		return ""
	}

	// TODO: remove this when the issue https://github.com/golang/go/issues/29228 is fixed
	if bi.Main.Version == "(devel)" || bi.Main.Version == "" {
		return ""
	}

	return bi.Main.Version
}

func getCommit(bi *debug.BuildInfo) string {
	return getKey(bi, "vcs.revision")
}

func getBuildDate(bi *debug.BuildInfo) string {
	buildTime := getKey(bi, "vcs.time")
	t, err := time.Parse("2006-01-02T15:04:05Z", buildTime)
	if err != nil {
		return ""
	}
	return t.Format("2006-01-02T15:04:05")
}

func getKey(bi *debug.BuildInfo, key string) string {
	if bi == nil {
		return ""
	}
	for _, iter := range bi.Settings {
		if iter.Key == key {
			return iter.Value
		}
	}
	return ""
}

func firstNonEmpty(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}

func GetVersion(options ...Option) Version {
	buildInfo := getBuildInfo()
	v := Version{
		Version:   firstNonEmpty(getGitVersion(buildInfo), "devel"),
		GitCommit: firstNonEmpty(getCommit(buildInfo), unknown),
		BuildDate: firstNonEmpty(getBuildDate(buildInfo), unknown),
		BuiltBy:   unknown,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}

	for _, opt := range options {
		opt(&v)
	}

	return v
}

func (v Version) String() string {
	b := strings.Builder{}
	w := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)

	if v.AsciiArt != "" {
		_, _ = fmt.Fprint(w, v.AsciiArt)
	}

	_, _ = fmt.Fprintf(w, "Version:\t%s\n", v.Version)
	_, _ = fmt.Fprintf(w, "GitCommit:\t%s\n", v.GitCommit)
	_, _ = fmt.Fprintf(w, "BuildDate:\t%s\n", v.BuildDate)
	_, _ = fmt.Fprintf(w, "BuiltBy:\t%s\n", v.BuiltBy)
	_, _ = fmt.Fprintf(w, "GoVersion:\t%s\n", v.GoVersion)
	_, _ = fmt.Fprintf(w, "Compiler:\t%s\n", v.Compiler)
	_, _ = fmt.Fprintf(w, "Platform:\t%s\n", v.Platform)

	_ = w.Flush()
	return b.String()
}

func (v *Version) JSONString() (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
