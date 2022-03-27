package cache

import (
	"context"
)

type cacheKey struct{}

// FromContext returns cache from context
func FromContext(ctx context.Context) (Cache, bool) {
	c, ok := ctx.Value(cacheKey{}).(Cache)
	return c, ok
}

// NewContext returns new context with cache
func NewContext(ctx context.Context, c Cache) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, cacheKey{}, c)
}
