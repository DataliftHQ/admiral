package config

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"runtime"

	"golang.org/x/oauth2"
)

type Config struct {
	path string `json:"-"`

	Settings Settings     `json:"settings"`
	Token    oauth2.Token `json:"token"`
}

type Settings struct {
	ServerAddress string `json:"server"`
	PlainText     bool   `json:"plainText"`
	Insecure      bool   `json:"insecure"`
	CertPEMData   []byte
	ClientCert    *tls.Certificate // TODO: THINK ABOUT THIS
	UserAgent     string           `json:"userAgent,omitempty"`
	Headers       []string         `json:"headers,omitempty"`
}

func DefaultPath() (string, error) {
	if configFile := os.Getenv("ADMIRAL_CONFIG_FILE"); configFile != "" {
		return configFile, nil
	}

	if xdgConfigHome := os.Getenv("XDG_CONFIG_HOME"); xdgConfigHome != "" {
		return path.Join(xdgConfigHome, "admiral", "admiral.json"), nil
	}

	if winConfigHome := os.Getenv("AppData"); winConfigHome != "" && runtime.GOOS == "windows" {
		return path.Join(winConfigHome, "Admiral", "admiral.json"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(homeDir, ".config", "admiral", "admiral.json"), nil
}

func Read(path string) (*Config, error) {
	var err error
	var config Config

	if fi, err := os.Stat(path); err == nil {
		err = getFilePermission(fi)
		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return &Config{path: path}, nil
		} else {
			return nil, err
		}
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	// set path after unmarshalling
	config.path = path

	return &config, nil
}

func (c *Config) Save() error {
	err := os.MkdirAll(path.Dir(c.path), os.ModePerm)
	if err != nil {
		return err
	}

	data, err := json.Marshal(c)
	if err == nil {
		err = os.WriteFile(c.path, data, 0600)
	}
	return err
}

func (c *Config) Delete() error {
	_, err := os.Stat(c.path)
	if errors.Is(err, fs.ErrNotExist) {
		return err
	}
	return os.Remove(c.path)
}

func getFilePermission(fi os.FileInfo) error {
	if fi.Mode().Perm() == 0600 || fi.Mode().Perm() == 0400 {
		return nil
	}
	return fmt.Errorf("config file has incorrect permission flags:%s."+
		"change the file permission either to 0400 or 0600.", fi.Mode().Perm().String())
}
