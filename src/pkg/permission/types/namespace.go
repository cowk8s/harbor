package types

import "sync"

var (
	parsesMu sync.RWMutex
	parses   = map[string]NamespaceParse{}
)

// NamespaceParse parse namespace from the resource
type NamespaceParse func(Resource) (Namespace, bool)

// Namespace the namespace interface
type Namespace interface {
	// Kind returns the kind of namespace
	Kind() string
	// Resource returns new resource for subresources with the namespace
	Resource(subresource ...Resource) Resource
	// Identity returns identity attached with namespace
	Identity() interface{}
	// GetPolicies returns all policies of the namespace
	GetPolicies() []*Policy
}
