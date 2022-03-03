package middleware

import "net/http"

// Middleware receives a handler and returns another handler.
// The returned handler can do some customized task according to
// the requirement
type Middleware func(http.Handler) http.Handler

// Chain make middlewares together
func Chain(middlewares ...Middleware) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			h = middlewares[i](h)
		}

		return h
	}
}

// WithMiddlewares apply the middlewares to the handler.
// The middlewares are executed in the order that they are applied
func WithMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	return Chain(middlewares...)(handler)
}

func New(fn func(http.ResponseWriter, *http.Request, http.Handler), skippers ...Skipper) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, skipper := range skippers {
				if skipper(r) {
					next.ServeHTTP(w, r)
					return
				}
			}

			fn(w, r, next)
		})
	}
}

func BeforeRequest(hook func(*http.Request) error, skippers ...Skipper) func(http.Handler) http.Handler {
	return New(func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		if err := hook(r); err != nil {
			return
		}

		next.ServeHTTP(w, r)
	}, skippers...)
}

func AfterResponse(hook func(http.ResponseWriter, *http.Request, int) error, skipper ...Skipper) func(http.Handler) http.Handler {
	return New(func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		res, ok := w.()
	})
}