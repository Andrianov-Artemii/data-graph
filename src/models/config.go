package models

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MediaAPIConfig MediaAPIConfig `yaml:"mediaConfing"`
}

func NewConfigYML(path string) (*Config, error) {
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

func NewConfigENV() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("MEDIA_SERVER_PORT"))
	if err != nil {
		return nil, err
	}
	return &Config{
		MediaAPIConfig: MediaAPIConfig{
			Host:            os.Getenv("MEDIA_SERVER_HOST"),
			Port:            port,
			AdminPass:       os.Getenv("MEDIA_SERVER_ADMIN_PASS"),
			StorageRootPath: os.Getenv("MEDIA_SERVER_STORAGE_ROOT_PATH"),
			Routes: MediaAPIRoutes{
				DataRoute: Router{
					Name:         os.Getenv("MEDIA_SERVER_DATA_ROUTE_NAME"),
					StorageRoute: os.Getenv("MEDIA_SERVER_DATA_ROUTE_STORAGE_ROUTE"),
				},
			},
		},
	}, nil
}

type MediaAPIConfig struct {
	Host            string         `yaml:"host"`
	Port            int            `yaml:"port"`
	AdminPass       string         `yaml:"adminPass"`
	StorageRootPath string         `yaml:"storageRootPath"`
	Routes          MediaAPIRoutes `yaml:"routes"`
}

type MediaAPIRoutes struct {
	DataRoute Router `yaml:"dataRoute"`
}

type Router struct {
	Name         string `yaml:"name"`
	StorageRoute string `yaml:"storageRoute"`
}
