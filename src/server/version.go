package server

import (
	"encoding/json"
	"net/http"

	lib_http "github.com/cowk8s/harbor/src/lib/http"

	"github.com/cowk8s/harbor/src/server/v2.0/route"
)

var (
	version = route.APIVersion
)

// APIVersion model
type APIVersion struct {
	Version string `json:"version"`
}

// GetAPIVersion returns the current supported API version
func GetAPIVersion(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&APIVersion{Version: version}); err != nil {
		lib_http.SendError(w, err)
	}
}
