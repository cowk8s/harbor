package router

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/astaxie/beego"
	beegocontext "github.com/astaxie/beego/context"
)

// ContextKeyInput ...
type ContextKeyInput struct{}

func NewRoute() *Route {
	return &Route{}
}

type Route struct {
	parent  *Route
	methods []string
	path    string
}

func (r *Route) Path(path string) *Route {
	r.path = path
	return r
}

func (r *Route) Handler(handler http.Handler) {
	methods := r.methods
	if len(methods) == 0 && r.parent != nil {
		methods = r.parent.methods
	}

	path := r.path
	if r.parent != nil {
		path = filepath.Join(r.parent.path, path)
	}

	filterFunc := beego.FilterFunc(func(ctx *beegocontext.Context) {
		ctx.Request = ctx.Request.WithContext(
			context.WithValue(ctx.Request.Context(), ContextKeyInput{}, ctx.Input))
	})

	if len(methods) == 0 {
		beego.Any(path, filterFunc)
		return
	}
	for _, method := range methods {
		switch method {
		case http.MethodGet:
			beego.Get(path, filterFunc)
		case http.MethodHead:
			beego.Head(path, filterFunc)
		case http.MethodPut:
			beego.Put(path, filterFunc)
		case http.MethodPatch:
			beego.Patch(path, filterFunc)
		case http.MethodPost:
			beego.Post(path, filterFunc)
		case http.MethodDelete:
			beego.Delete(path, filterFunc)
		case http.MethodOptions:
			beego.Options(path, filterFunc)
		}
	}
}
