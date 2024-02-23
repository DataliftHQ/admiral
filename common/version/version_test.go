package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithVersionOption(t *testing.T) {
	expectedVersion := "1.0.0"
	v := GetVersion(WithVersion(expectedVersion))

	assert.Equal(t, expectedVersion, v.Version, "expected version = %v, got %v", v.Version, expectedVersion)
}

func TestWithGitCommitOption(t *testing.T) {
	expectedCommit := "abcdefg"
	v := GetVersion(WithGitCommit(expectedCommit))

	assert.Equal(t, expectedCommit, v.GitCommit, "expected git commit = %v, got %v", v.GitCommit, expectedCommit)
}

func TestWithBuiltByOption(t *testing.T) {
	expectedBuiltBy := "devteam"
	v := GetVersion(WithBuiltBy(expectedBuiltBy))

	assert.Equal(t, expectedBuiltBy, v.BuiltBy, "expected built by = %v, got %v", v.BuiltBy, expectedBuiltBy)
}

func TestWithAsciiArtOption(t *testing.T) {
	expectedArt := "ASCII ART"
	v := GetVersion(WithAsciiArt(expectedArt))

	assert.Equal(t, expectedArt, v.AsciiArt, "expected ascii art = %v, got %v", v.AsciiArt, expectedArt)
}

func TestDefaultValues(t *testing.T) {
	v := GetVersion() // No options provided

	assert.Equal(t, "devel", v.Version, "expected version = %v, got devel", v.AsciiArt)
	assert.Equal(t, unknown, v.GitCommit, "expected git commit = %v, got %v", v.AsciiArt, unknown)
	assert.Equal(t, unknown, v.BuildDate, "expected build date = %v, got %v", v.AsciiArt, unknown)
	assert.Equal(t, unknown, v.BuiltBy, "expected built by = %v, got %v", v.BuiltBy, unknown)
	assert.Equal(t, "", v.AsciiArt, "expected ascii art = %v, got %v", v.AsciiArt, "")
}

func TestStringOutput(t *testing.T) {
	v := Version{
		Version:   "1.0.0",
		GitCommit: "abcdef",
		BuildDate: "2023-01-01T00:00:00",
		BuiltBy:   "devteam",
		GoVersion: "go1.15",
		Compiler:  "gc",
		Platform:  "linux/amd64",
	}

	expectedString := "Version:    1.0.0\nGitCommit:  abcdef\nBuildDate:  2023-01-01T00:00:00\nBuiltBy:    devteam\nGoVersion:  go1.15\nCompiler:   gc\nPlatform:   linux/amd64\n"
	assert.Equal(t, expectedString, v.String(), "expected text string = %v, got %v", v.String(), expectedString)
}

func TestJSONStringOutput(t *testing.T) {
	v := Version{
		Version:   "1.0.0",
		GitCommit: "abcdef",
		BuildDate: "2023-01-01T00:00:00",
		BuiltBy:   "devteam",
		GoVersion: "go1.15",
		Compiler:  "gc",
		Platform:  "linux/amd64",
	}

	jsonString, err := v.JSONString()
	if err != nil {
		t.Errorf("Version.JSONString() error = %v", err)
		return
	}

	expectedJsonString := `{
  "version": "1.0.0",
  "gitCommit": "abcdef",
  "buildDate": "2023-01-01T00:00:00",
  "builtBy": "devteam",
  "goVersion": "go1.15",
  "compiler": "gc",
  "platform": "linux/amd64"
}`

	assert.Equal(t, expectedJsonString, jsonString, "expected json string = %v, got %v", expectedJsonString, jsonString)
}
