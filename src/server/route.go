package server

import (
	"net/http"

	"github.com/cowk8s/harbor/src/server/router"
)

func registerRoutes() {
	// API version
	router.NewRoute().Method(http.MethodGet).Path("/api/version").HandlerFunc(GetAPIVersion)
}
