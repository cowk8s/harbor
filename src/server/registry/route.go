package registry

import (
	"net/http"

	"github.com/cowk8s/harbor/src/server/middleware/v2auth"
	"github.com/cowk8s/harbor/src/server/router"
)

func RegisterRoutes() {
	root := router.NewRoute().
		Path("/v2").
		Middleware(v2auth.Middleware())
	// catalog
	root.NewRoute().
		Method(http.MethodGet).
		Path("_catalog").
		Handler()
}
