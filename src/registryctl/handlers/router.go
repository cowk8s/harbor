package handlers

import (
	"net/http"

	"github.com/cowk8s/harbor/src/registryctl/config"
	"github.com/gorilla/mux"
)

func newRouter(conf config.Configuration) http.Handler {
	// create the root rooter
	rootRouter := mux.NewRouter()
	rootRouter.StrictSlash(true)
	rootRouter.HandleFunc("/api/health", api.Health).Methods("GET")

	rootRouter.Path("/api/registry/blob/{reference}").Methods(http.MethodDelete).Handler(blob)
	return rootRouter
}
