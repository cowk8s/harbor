package metadata

type Item struct {
	Name string `json:"name,omitempty"`
}

var (
	ConfigList = []Item{
		{Name: "hi"},
	}
)
