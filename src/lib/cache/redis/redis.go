package redis

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/cowk8s/harbor/src/lib/cache"
	"github.com/go-redis/redis/v8"
)

var _ cache.Cache = (*Ca)

// Cache redis cache
type Cache struct {
	*redis.Client
	opts *cache.Options
}

// Contains returns true if key exists
func (c *Cache) Contains(ctx context.Context, key string) bool {
	val, err := c.Client.Exists(ctx, c.opts.Key(key)).Result()
	if err != nil {
		return false
	}

	return val == 1
}

// Delete delete item from cache by key
func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.Client.Del(ctx, c.opts.Key(key)).Err()
}

// Fetch retrieve the cached key value
func (c *Cache) Fetch(ctx context.Context, key string, value interface{}) error {
	data, err := c.Client.Get(ctx, c.opts.Key(key)).Bytes()
	if err != nil {
		// convert internal or Timeout error to be ErrNotFound
		// so that the caller can continue working without breaking
		// return cache.ErrNotFound
		return fmt.Errorf("%w:%v", cache.ErrNotFound, err)
	}

	if err := c.opts.Codec.Decode(data, value); err != nil {
		return errors.Errorf("failed to decode cached value to dest, key %s, error: %v", key, err)
	}

	return nil
}

func New(opts cache.Options) (cache.Cache, error) {
	if opts.Address == "" {
		opts.Address = "redis://localhost:6379/0"
	}

	if opts.Codec == nil {
		opts.Codec = cache.DefaultCodec()
	}

	u, err := url.Parse(opts.Address)
	if err != nil {
		return nil, err
	}

	// For compatibility, should convert idle_timeout_seconds to idle_timeout.
	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	if t := values.Get("idle_timeout_seconds"); t != "" {
		values.Del("idle_timeout_seconds")
		values.Set("idle_timeout", t)
		u.RawQuery = values.Encode()
	}

	var client *redis.Client

	switch u.Scheme {
	case cache.Redis:
		rdbOpts, err := redis.ParseURL(u.String())
		if err != nil {
			return nil, err
		}

		client = redis.NewClient(rdbOpts)
	case cache.RedisSentinel:
		failoverOpts, err := 
	}
}

func init() {
	cache.Register(cache.Redis, New)
	cache.Register(cache.RedisSentinel, New)
}
