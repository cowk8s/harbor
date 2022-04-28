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

// NewRoute creates a new route
func NewRoute() *Route {
	return &Route{}
}

// Route stores the information that matches a request
type Route struct {
	parent  *Route
	methods []string
	path    string
}

// NewRoute returns a sub route based on the current one
func (r *Route) NewRoute() *Route {
	return &Route{
		parent: r,
	}
}

// Method sets the method that the route matches
func (r *Route) Method(method string) *Route {
	r.methods = append(r.methods, method)
	return r
}

// Path sets the path that the route matches. Path uses the beego router path pattern
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

	for _, method := range methods {
		switch method {
		case http.MethodGet:
			beego.Get(path, filterFunc)
		}
	}
}

// HandlerFunc sets the handler function that handles the request
func (r *Route) HandlerFunc(f http.HandlerFunc) {
	r.Handler(f)
}
