package accessory

var (
	// Mgr is a global artifact manager instance
	Mgr = NewManager()
)

// Managaer is the only interface of artifact module to provide the management functions for artifacts
type Managaer interface {
	
}