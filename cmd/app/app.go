package main

import (
	"flag"
	"fmt"
	"log"
	"media-server/internal/api"
	"media-server/internal/domain"
)

type MODE string

const (
	MODE_ENV  MODE = "env"
	MODE_YAML MODE = "yaml"
)

func main() {
	modeFlag := flag.String("mode", "env", "Mode of configs loadings")
	configFlag := flag.String("config", "config.yaml", "Config destination (.yml)")
	flag.Parse()
	var (
		config *domain.Config
		err    error
	)
	switch *modeFlag {
	case string(MODE_YAML):
		log.Println("Usings .yaml file for configuration")
		if configFlag == nil {
			log.Panic("Config folder is not setted in arguments!")
		}
		if config, err = domain.NewConfigYML(*configFlag); err != nil {
			log.Panic(fmt.Sprintf("Error while configs reading!\nError: %s", err.Error()))
		}
	case string(MODE_ENV):
		log.Println("Loading configuration from ENV")
		if config, err = domain.NewConfigENV(); err != nil {
			log.Panic(fmt.Sprintf("Error while configs readings from envs!\nError: %s", err.Error()))
		}
	}
	api, _ := api.NewMediaAPI(&config.MediaAPIConfig)
	api.Run()
}
