package exporter

import (
	"sync"
	"time"
)

var c *cache

type cachedValue struct {
	Value      interface{}
	Expiration int64
}

type cache struct {
	CacheDuration int64
	store         map[string]cachedValue
	*sync.RWMutex
}

func CacheGet(key string) (value interface{}, ok bool) {
	c.RLock()
	v, ok := c.store[key]
	c.RUnlock()
	if !ok {
		return nil, false
	}
	if time.Now().Unix() > v.Expiration {
		c.Lock()
		delete(c.store, key)
		c.RUnlock()
		return nil, false
	}
	return v.Value, true
}
