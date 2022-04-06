package types

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

// Resource the type of resource
type Resource string

func (res Resource) String() string {
	return string(res)
}

// RelativeTo returns relative resource to other resource
func (res Resource) RelativeTo(other Resource) (Resource, error) {
	prefix := other.String()
	str := res.String()

	if !strings.HasPrefix(str, prefix) {
		return Resource(""), errors.New("value error")
	}

	relative := strings.TrimPrefix(strings.TrimPrefix(str, prefix), "/")
	if relative == "" {
		relative = "."
	}

	return Resource(relative), nil
}

// Subresource returns subresource
func (res Resource) Subresource(resources ...Resource) Resource {
	elements := []string{res.String()}

	for _, resource := range resources {
		elements = append(elements, resource.String())
	}

	return Resource(path.Join(elements...))
}

// GetNamespace returns namespace from resource
func (res Resource) GetNamespace() (Namespace, error) {
	return nil, fmt.Errorf("no namespace found for %s", res)
}
