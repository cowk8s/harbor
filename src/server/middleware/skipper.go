package middleware

import (
	"net/http"
	"path"
	"regexp"
)

// Skipper defines a function to skip middleware.
// Returning true skips processing the middleware.
type Skipper func(*http.Request) bool

// MethodAndPathSkipper returns skipper which
// will skip the middleware when r.Method equals the method and r.URL.Path matches the re
// when method is "*" it equals all http method
func MethodAndPathSkipper(method string, re *regexp.Regexp) func(r *http.Request) bool {
	return func(r *http.Request) bool {
		path := path.Clean(r.URL.EscapedPath())
		if (method == "*" || r.Method == method) && re.MatchString(path) {
			return true
		}

		return false
	}
}

// NegativeSkipper returns skipper which is negative of the input skipper
func NegativeSkipper(skipper Skipper) func(*http.Request) bool {
	return func(r *http.Request) bool {
		return !skipper(r)
	}
}
