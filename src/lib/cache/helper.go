package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	fetchOrSaveMu = keyMutex{m: &sync.Map{}}
)

// FetchOrSave retrieves the value for the key if present in the cache.
// Otherwise, it saves the value from the builder and retrieves the value for the key again.
func FetchOrSave(ctx context.Context, c Cache, key string, value interface{}, builder func() (interface{}, error), expiration ...time.Duration) error {
	err := c.Fetch(ctx, key, value)
	// value found from the cache
	if err == nil {
		return nil
	}
	// internal error
	if !errors.Is(err, ErrNotFound) {
		return err
	}

	// lock the key in cache and try to build the value for the key
	lockKey := fmt.Sprintf("%p:%s", c, key)
	fetchOrSaveMu.Lock(lockKey)

	defer fetchOrSaveMu.Unlock(lockKey)

	// fetch again to avoid to build the value multi-times
	err = c.Fetch(ctx, key, value)
	if err == nil {
		return nil
	}
	// internal error
	if !errors.Is(err, ErrNotFound) {
		return err
	}

	return c.Fetch(ctx, key, value) // after the building, fetch value again
}
