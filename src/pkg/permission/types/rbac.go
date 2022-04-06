package types

type RBACRole interface {
	// GetRoleName returns the role identity, if empty string role's policies will be ignore
	GetRoleName() string
	// GetPolicies returns the policies of the role
	GetPolicies() []*Policy
}

type RBACUser interface {
	GetUserName() string

	GetPolicies() []*Policy

	GetRoles() []RBACRole
}
