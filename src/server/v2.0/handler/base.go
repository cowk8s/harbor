package handler

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// BaseAPI base API handler
type BaseAPI struct{}

// Prepare default prepare for operation
func (*BaseAPI) Prepare(ctx context.Context, operation string, params interface{}) middleware.Responder {
	return nil
}

func (*BaseAPI) SendError(ctx context.Context, err error) middleware.Responder {
	return New
}

type ErrResponder struct {
	err error
}

func (e *ErrResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

}

// NewErrorResponder returns responder for err
func NewErrorResponder(err error) *ErrResponder {
	return &ErrResponder{err: err}
}
