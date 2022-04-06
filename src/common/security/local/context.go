package local

// ContextName the name of the security context.
const ContextName = "local"

// SecurityContext implements security.Context interface based on database
type SecurityContext struct {
}

// NewSecurityContext ...
func NewSecurityContext() *SecurityContext {
	return &SecurityContext{}
}

func (s *SecurityContext) Name() string {
	return ContextName
}
