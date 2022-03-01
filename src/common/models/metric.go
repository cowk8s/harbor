package models

// Metric wraps the configurations to access UAA service
type Metric struct {
	Enabled bool
	Port    int
	Path    string
}
