package models

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MediaAPIConfig MediaAPIConfig `yaml:"mediaConfing"`
}

func NewConfig(path string) (*Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	yamlDecoder := yaml.NewDecoder(configFile)
	var config Config
	err = yamlDecoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

type MediaAPIConfig struct {
	Host            string         `yaml:"host"`
	Port            int            `yaml:"port"`
	StorageRootPath string         `yaml:"storageRootPath"`
	Routes          MediaAPIRoutes `yaml:"routes"`
}

type MediaAPIRoutes struct {
	ImageRoute Router `yaml:"imageRoute"`
	VideoRoute Router `yaml:"videoRoute"`
}

type Router struct {
	Name         string `yaml:"name"`
	StorageRoute string `yaml:"storageRoute"`
}
