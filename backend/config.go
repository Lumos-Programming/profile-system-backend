package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port      int             `yaml:"port"`
	Firestore FirestoreConfig `yaml:"firestore"`
}

type FirestoreConfig struct {
	ProjectID   string `yaml:"project_id"`
	Credentials string `yaml:"credentials"`
}

const configPath = "../secrets/config.yaml"

func loadConfig() (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
