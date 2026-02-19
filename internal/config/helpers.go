package config

import (
	"encoding/json"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"
const defaultFilePerm = 0666

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return homeDir, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	configFile := path.Join(configPath, configFileName)
	jsonBytes, err := json.Marshal(cfg)

	if err != nil {
		return err
	}
	return os.WriteFile(configFile, jsonBytes, defaultFilePerm)
}
