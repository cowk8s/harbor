package lib

import (
	"context"
)

type contextKey string

// define all context key here to avoid conflict
const (
	contextKeyAPIVersion   contextKey = "apiVersion"
	contextKeyArtifactInfo contextKey = "artifactInfo"
	contextKeyAuthMode     contextKey = "authMode"
	contextKeyCarrySession contextKey = "carrySession"
)

// ArtifactInfo wraps the artifact info extracted from the request to "/v2/"
type ArtifactInfo struct {
	Repository           string
	Reference            string
	ProjectName          string
	Digest               string
	Tag                  string
	BlobMountRepository  string
	BlobMountProjectName string
	BlobMountDigest      string
}

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

// WithAPIVersion returns a context with APIVersion set
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

// WithArtifactInfo returns a context with ArtifactInfo set
func WithArtifactInfo(ctx context.Context, art ArtifactInfo) context.Context {
	return setToContext(ctx, contextKeyArtifactInfo, art)
}

// GetArtifactInfo gets the ArtifactInfo from the context
func GetArtifactInfo(ctx context.Context) (art ArtifactInfo) {
	value := getFromContext(ctx, contextKeyArtifactInfo)
	if value != nil {
		art, _ = value.(ArtifactInfo)
	}
	return
}

// WithAuthMode returns a context with auth mode set
func WithAuthMode(ctx context.Context, mode string) context.Context {
	return setToContext(ctx, contextKeyAuthMode, mode)
}

// GetAuthMode gets the auth mode from the context
func GetAuthMode(ctx context.Context) string {
	mode := ""
	value := getFromContext(ctx, contextKeyAuthMode)
	if value != nil {
		mode, _ = value.(string)
	}
	return mode
}

// WithCarrySession returns a context with "carry session" set that indicates whether the request carries session or not
func WithCarrySession(ctx context.Context, carrySession bool) context.Context {
	return setToContext(ctx, contextKeyCarrySession, carrySession)
}

// GetCarrySession gets the "carry session" from the context indicates whether the request carries session or not
func GetCarrySession(ctx context.Context) bool {
	carrySession := false
	value := getFromContext(ctx, contextKeyCarrySession)
	if value != nil {
		carrySession, _ = value.(bool)
	}
	return carrySession
}
