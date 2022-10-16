package main

import (
	"flag"
	"fmt"
	"log"
	"media-server/src/api"
	"media-server/src/models"
)

func main() {
	configFlag := flag.String("config", "config.yaml", "a string")
	flag.Parse()
	if configFlag == nil {
		log.Panic("Config folder is not setted in arguments!")
	}
	config, err := models.NewConfig(*configFlag)
	if err != nil {
		log.Panic(fmt.Sprintf("Error while configs reading!\nError: %s", err.Error()))
	}
	api, _ := api.NewMediaAPI(&config.MediaAPIConfig)
	api.Run()
}
