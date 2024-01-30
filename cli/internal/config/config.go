package config

import (
	"io"
	"os"
	"path"
	"runtime"
	"sync"

	"gopkg.in/yaml.v2"
)

var saveMutex sync.Mutex

type Config struct {
	Version       int    `yaml:"version,omitempty" json:"version,omitempty"`
	ServerAddress string `yaml:"server_address,omitempty" json:"server_address,omitempty"`
	OAuth2        OAuth2 `yaml:"oauth2,omitempty" json:"oauth2,omitempty"`
}

type OAuth2 struct {
	Issuer   string   `yaml:"issuer,omitempty" json:"issuer,omitempty"`
	ClientId string   `yaml:"client_id,omitempty" json:"client_id,omitempty"`
	Scopes   []string `yaml:"scopes,omitempty" json:"scopes,omitempty"`
}

func Load(file string) (config Config, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	return LoadReader(f)
}

func LoadReader(fd io.Reader) (config Config, err error) {
	data, err := io.ReadAll(fd)
	if err != nil {
		return config, err
	}

	err = yaml.UnmarshalStrict(data, &config)
	return config, err
}

func Save(file string, config Config) error {
	saveMutex.Lock()
	defer saveMutex.Unlock()

	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0600)
}

func ConfigDir() (string, error) {
	if configDir := os.Getenv("ADMIRAL_CONFIG_DIR"); configDir != "" {
		return configDir, nil
	}

	if xdgConfigHome := os.Getenv("XDG_CONFIG_HOME"); xdgConfigHome != "" {
		return path.Join(xdgConfigHome, "admiral"), nil
	}

	if winConfigHome := os.Getenv("AppData"); winConfigHome != "" && runtime.GOOS == "windows" {
		return path.Join(winConfigHome, "Admiral"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(homeDir, ".config", "admiral"), nil
}
