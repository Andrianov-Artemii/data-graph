package apiServer

import (
	"data-graph-backend/pkg/dbConnector"
	"data-graph-backend/pkg/properties"
	"log"
	"net/http"
)

func Start(config *properties.Config) error {
	dbConnection, err := dbConnector.NewConnection(config)
	if err != nil {
		log.Fatal("Can't connect to db - ", err.Error())
	}

	router := &Router{
		dbConnector: dbConnection,
	}

	configureRouters(router)

	log.Println("Listening on " + config.ProgramSettings.Host + config.ProgramSettings.Port)
	return http.ListenAndServe(config.ProgramSettings.Port, nil)
}
