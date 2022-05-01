package api

import (
	"context"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	defaultPageSize int64 = 500
	maxPageSize     int64 = 500

	// APIVersion is the current core api version
	APIVersion = "v2.0"
)

// BaseAPI wraps common methods for controllers to host API
type BaseAPI struct {
	beego.Controller
}

// Context returns the context.Context from http.Request
func (b *BaseAPI) Context() context.Context {
	return b.Ctx.Request.Context()
}

// GetStringFromPath gets the param from path and returns it as string
func (b *BaseAPI) GetStringFromPath(key string) string {
	return b.Ctx.Input.Param(key)
}

// GetInt64FromPath gets the param from path and returns it as int64
func (b *BaseAPI) GetInt64FromPath(key string) (int64, error) {
	value := b.Ctx.Input.Param(key)
	return strconv.ParseInt(value, 10, 64)
}
