package types

const (
	// EffectAllow allow effect
	EffectAllow = Effect("allow")
	// EffectDeny deny effect
	EffectDeny = Effect("deny")
)

// Action the type of action
type Action string

func (act Action) String() string {
	return string(act)
}

// Effect the type of effect
type Effect string

func (eff Effect) String() string {
	return string(eff)
}

// Policy the type of policy
type Policy struct {
	Resource `json:"resource,omitempty"`
	Action   `json:"action,omitempty"`
	Effect   `json:"effect,omitempty"`
}

// GetEffect returns effect of resource, default is allow
func (p *Policy) GetEffect() string {
	eft := p.Effect
	if eft == "" {
		eft = EffectAllow
	}

	return eft.String()
}

func (p *Policy) String() string {
	return p.Resource.String() + ":" + p.Action.String() + ":" + p.GetEffect()
}
