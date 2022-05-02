package blob

import (
	"net/http"

	storagedriver "github.com/docker/distribution/registry/storage/driver"
)

const tracerName = "goharbor/harbor/src/registryctl/api/registry/blob"

// NewHandler returns the handler to handler blob request
func NewHandler(storageDriver storagedriver.StorageDriver) http.Handler {
	return &handler{
		storageDriver: storageDriver,
	}
}

type handler struct {
	storageDriver storagedriver.StorageDriver
}

// ServeHTTP ...
func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodDelete:
		h.delete(w, req)
	default:
		api.HandlerNotMethodAllowed(w)
	}
}

// DeleteBlob ...
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {

}
