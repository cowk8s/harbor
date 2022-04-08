package lib

type void = struct{}

// Set a simple set
type Set map[interface{}]void

// Add add item to set
func (s Set) Add(item interface{}) {
	s[item] = void{}
}

// Exists returns true when item in the set
func (s Set) Exists(item interface{}) bool {
	_, ok := s[item]

	return ok
}

// Items returns the items in the set
func (s Set) Items() []interface{} {
	var items []interface{}
	for item := range s {
		items = append(items, item)
	}

	return items
}
