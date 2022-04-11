package selector

import (
	"encoding/json"

	"github.com/cowk8s/harbor/src/lib/errors"
)

const (
	// Image kind
	Image = "image"
	// Chart kind
	Chart = "chart"
)

// Repository of candidate
type Repository struct {
	// Namespace(project) ID
	NamespaceID int64
	// Namespace
	Namespace string `json:"namespace"`
	// Repository name
	Name string `json:"name"`
	// So far we need the kind of repository and retrieve candidates with different APIs
	// TODO: REMOVE IT IN THE FUTURE IF WE SUPPORT UNIFIED ARTIFACT MODEL
	Kind string `json:"kind"`
}

// ToJSON marshals repository to JSON string
func (r *Repository) ToJSON() (string, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return "", errors.Wrap(err, "marshal reporitory")
	}

	return string(jsonData), nil
}