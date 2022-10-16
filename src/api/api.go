package api

import (
	"fmt"
	"log"
	"media-server/src/models"
	"net/http"
)

type MediaAPI struct {
	Host     string
	Port     int
	RootPath string
	Routes   models.MediaAPIRoutes
}

func NewMediaAPI(config *models.MediaAPIConfig) (*MediaAPI, error) {
	api := &MediaAPI{Host: config.Host, Port: config.Port, RootPath: config.StorageRootPath, Routes: config.Routes}
	return api, nil
}

func (api *MediaAPI) Run() {
	http.HandleFunc(api.Routes.ImageRoute.Name, api.getImageByUrl)
	http.HandleFunc(api.Routes.VideoRoute.Name, api.getVideoByUrl)
	log.Printf("Run server on %s:%d", api.Host, api.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", api.Host, api.Port), nil)
}

func (api *MediaAPI) getImageByUrl(rw http.ResponseWriter, r *http.Request) {
	currentPath := r.URL.RequestURI()[len(api.Routes.ImageRoute.Name):]
	http.ServeFile(rw, r, fmt.Sprintf("%s%s%s", api.RootPath, api.Routes.ImageRoute.StorageRoute, currentPath))
}

func (api *MediaAPI) getVideoByUrl(rw http.ResponseWriter, r *http.Request) {
	currentPath := r.URL.RequestURI()[len(api.Routes.VideoRoute.Name):]
	http.ServeFile(rw, r, fmt.Sprintf("%s%s%s", api.RootPath, api.Routes.VideoRoute.StorageRoute, currentPath))
}
