package config

import (
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
	Settings Settings      `json:"settings"`
	Token    *oauth2.Token `json:"token"`
}

type Settings struct {
	ServerAddress string `json:"server"`
	PlainText     bool   `json:"plainText"`
	Insecure      bool   `json:"insecure"`
	CertPEMData   []byte
	ClientCert    []byte
}

func DefaultFile() (string, error) {
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

func ReadFile(path string) (*Config, error) {
	return nil, nil
}

func ReadConfig(path string) (*Config, error) {
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
			return nil, nil
		} else {
			return nil, err
		}
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func WriteConfig(config Config, configPath string) error {
	err := os.MkdirAll(path.Dir(configPath), os.ModePerm)
	if err != nil {
		return err
	}

	data, err := json.Marshal(config)
	if err == nil {
		err = os.WriteFile(configPath, data, 0600)
	}
	return err
}

func DeleteConfig(configPath string) error {
	_, err := os.Stat(configPath)
	if errors.Is(err, fs.ErrNotExist) {
		return err
	}
	return os.Remove(configPath)
}

func getFilePermission(fi os.FileInfo) error {
	if fi.Mode().Perm() == 0600 || fi.Mode().Perm() == 0400 {
		return nil
	}
	return fmt.Errorf("config file has incorrect permission flags:%s."+
		"change the file permission either to 0400 or 0600.", fi.Mode().Perm().String())
}
