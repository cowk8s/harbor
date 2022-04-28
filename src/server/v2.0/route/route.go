package route

import (
	"github.com/cowk8s/harbor/src/server/router"
	"github.com/cowk8s/harbor/src/server/v2.0/handler"
)

// const definition
const (
	APIVersion = "v2.0"
)

// RegisterRoutes for Harbor v2.0 APIs
func RegisterRoutes() {
	router.NewRoute().Path("/api/" + APIVersion + "/*").
		Handler(handler.New())
}
