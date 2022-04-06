package lib

import (
	"context"
)

type contextKey string

const (
	contextKeyAPIVersion contextKey = "apiVersion"
)

func setToContext(ctx context.Context, key contextKey, value interface{}) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, key, value)
}

func getFromContext(ctx context.Context, key contextKey) interface{} {
	if ctx == nil {
		return nil
	}
	return ctx.Value(key)
}

func WithAPIVersion(ctx context.Context, version string) context.Context {
	return setToContext(ctx, contextKeyAPIVersion, version)
}

// GetAPIVersion gets the API version from the context
func GetAPIVersion(ctx context.Context) string {
	version := ""
	value := getFromContext(ctx, contextKeyAPIVersion)
	if value != nil {
		version, _ = value.(string)
	}
	return version
}
