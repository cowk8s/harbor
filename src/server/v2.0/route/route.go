package route

import (
	"github.com/cowk8s/harbor/src/server/router"
	"github.com/cowk8s/harbor/src/server/v2.0/handler"
)

const (
	APIVersion = "v2.0"
)

func RegisterRoutes() {
	router.NewRoute().Path("/api/" + APIVersion + "/*").Handler(handler.New())
}
