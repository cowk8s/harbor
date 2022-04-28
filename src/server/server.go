package server

import (
	v2 "github.com/cowk8s/harbor/src/server/v2.0/route"
)

// RegisterRoutes register all routes
func RegisterRoutes() {
	registerRoutes()    // service/internal API/UI controller/etc.
	v2.RegisterRoutes() // v2.0 APIs
}
