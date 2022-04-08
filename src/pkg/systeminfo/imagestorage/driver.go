package imagestorage

// GlobalDriver is a global image storage driver
var GlobalDriver Driver

// Capacity holds information about capacity of image storage
type Capacity struct {
	// total size(byte)
	Total uint64 `json:"total"`
	// available size(byte)
	Free uint64 `json:"free"`
}

// Driver defines methods that an image storage driver must implement
type Driver interface {
	// Name returns a human-readable name of the driver
	Name() string
	// Cap returns the capacity of the image storage
	Cap() (*Capacity, error)
}
