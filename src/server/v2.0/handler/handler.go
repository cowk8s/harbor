package handler

import (
	"log"
	"net/http"

	"github.com/cowk8s/harbor/src/server/v2.0/restapi"
)

func New() http.Handler {
	h, _, err := restapi.HandlerAPI(restapi.Config{
		HealthAPI: newHealthAPI(),
	})
	if err != nil {
		log.Fatal(err)
	}

	return h
}
