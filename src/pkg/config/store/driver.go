// Package store is only used in the internal implement of manager, not a public api.
package store

import "context"

// Driver the interface to save/load config
type Driver interface {
	// Load - load config item from config driver
	Load(ctx context.Context) (map[string]interface{}, error)
	// Save - save config item into config driver
	Save(ctx context.Context, cfg map[string]interface{}) error
}