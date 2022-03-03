package handler

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
)

// BaseAPI base API handler
type BaseAPI struct{}

func (*BaseAPI) Prepare(ctx context.Context, operation string, params interface{}) middleware.Responder {
	return nil
}
