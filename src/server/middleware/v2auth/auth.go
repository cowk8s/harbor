package v2auth

import "net/http"

const (
	authHeader = "Authorization"
)

type reqChecker struct {
}

func (rc *reqChecker) check(req *http.Request) (string, error) {
	return "", nil
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(rw, req)
		})
	}
}
