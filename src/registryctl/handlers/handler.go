package handlers

import (
	"net/http"
	"os"

	"github.com/cowk8s/harbor/src/registryctl/config"
)

// NewHandlerChain returns a gorilla router which is wrapped by  authenticate handler
// and logging handler
func NewHandlerChain(conf config.Configuration) http.Handler {
	h := newRouter(conf)
	secrets := map[string]string{
		"jobSecret": os.Getenv("JOBSERVICE_SECRET"),
	}
	insecureAPIs := map[string]bool{
		"/api/health": true,
	}
	h =
}

type authHandler struct {
	authenicator auth
}