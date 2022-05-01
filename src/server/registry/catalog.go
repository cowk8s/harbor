package registry

import "net/http"

func newRepositoryHandler() http.Handler {
	return &repositoryHandler{}
}

type repositoryHandler struct {
}

func (r *repositoryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	return
}
