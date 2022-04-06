package security

import "context"

// Context abstracts the operations related with authN and authZ
type Context interface {
	// Name returns the name of the security context
	Name() string
	// IsAuthenticated returns whether the context has been authenticated or not
	IsAuthenticated() bool
	// GetUsername returns whether the user related to the context
	GetUsername() string
	// IsSysAdmin returns whether the user is system admin
	IsSysAdmin() bool
	// IsSolutionUser returns whether the user is solution user
	IsSolutionUser() bool
	// Can returns whether the user can do action on resource
	Can(ctx context.Context) bool
}

type securityKey struct{}

// NewContext returns context with security context
func NewContext(ctx context.Context, security Context) context.Context {
	return context.WithValue(ctx, securityKey{}, security)
}

// FromContext retruns security context from the context
func FromContext(ctx context.Context) (Context, bool) {
	c, ok := ctx.Value(securityKey{}).(Context)
	return c, ok
}
