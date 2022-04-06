package apiversion

import (
	"net/http"

	"github.com/cowk8s/harbor/src/lib"
	"github.com/cowk8s/harbor/src/server/middleware"
)

// Middleware returns a middleware that set the API version into the context
func Middleware(version string) middleware.Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx := lib.WithAPIVersion(req.Context(), version)
			handler.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
