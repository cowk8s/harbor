package lib

import "fmt"

// Link defines the model that describes the HTTP link header
type Link struct {
	URL   string
	Rel   string
	Attrs map[string]string
}

// String returns the string representation of a link
func (l *Link) String() string {
	s := fmt.Sprintf("<%s>", l.URL)
	if len(l.Rel) > 0 {
		s = fmt.Sprintf(`%s; rel="%s"`, s, l.Rel)
	}
	for key, value := range l.Attrs {
		s = fmt.Sprintf(`%s; %s="%s"`, s, key, value)
	}
	return s
}

// Links is a link object array
type Links []*Link

// String returns the string representation of links
func (l Links) String() string {

}
