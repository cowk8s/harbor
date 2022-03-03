package handler

import (
	"context"

	"github.com/cowk8s/harbor/src/server/v2.0/models"
	operations "github.com/cowk8s/harbor/src/server/v2.0/restapi/operations/health"
	"github.com/go-openapi/runtime/middleware"
)

func newHealthAPI() *healthAPI {
	return &healthAPI{}
}

type healthAPI struct {
	BaseAPI
}

func (r *healthAPI) GetHealth(ctx context.Context, params operations.GetHealthParams) middleware.Responder {
	s := &models.OverallHealthStatus{
		Status: "status",
	}
	return operations.NewGetHealthOK().WithPayload(s)
}
