package handler

import (
	"net/http"

	"github.com/cowk8s/harbor/src/server/v2.0/restapi"
)

// New returns http handler for API V2.0
func New() http.Handler {
	h, _, err := restapi.HandlerAPI(restapi.Config{})
	if err != nil {

	}

	return h
}
