package api

import (
	"net/http"

	"github.com/cowk8s/harbor/src/lib/log"
)

// Health ...
func Health(w http.ResponseWriter, r *http.Request) {
	if err := WriteJSON(w, "healthy"); err != nil {
		log.Errorf("Failed to write response: %v", err)
		return
	}
}
