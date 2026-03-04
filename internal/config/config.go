package config

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func Read() (Config, error) {
	config := Config{}
	configPath, err := getConfigFilePath()
	if err != nil {
		return config, err
	}
	configFile := path.Join(configPath, configFileName)

	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUsername = username

	return write(*cfg)
}
