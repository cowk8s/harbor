package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type middlewareTestSuite struct {
	suite.Suite
}

func (m *middlewareTestSuite) TestWithMiddlewares() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("key", w.Header().Get("key")+"handler")
	})

	middleware1 := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("key", w.Header().Get("key")+"middleware1")
			h.ServeHTTP(w, r)
		})
	}
	middleware2 := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("key", w.Header().Get("key")+"middleware2")
			h.ServeHTTP(w, r)
		})
	}

	record := &httptest.ResponseRecorder{}
	WithMiddlewares(handler, middleware1, middleware2).ServeHTTP(record, nil)
	m.Equal("middleware1middleware2handler", record.Header().Get("key"))
}

func TestMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, &middlewareTestSuite{})
}
