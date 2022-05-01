package v2auth

import (
	"context"
	"fmt"

	"github.com/cowk8s/harbor/src/common/rbac"
)

type target int

const (
	login target = iota
	catalog
	repository
)

func (t target) String() string {
	return []string{"login", "catalog", "repository"}[t]
}

type access struct {
	target target
	name   string
	action rbac.Action
}

func (a access) scopeStr(ctx context.Context) string {
	if a.target != repository {
		// Currently we do not support providing a token to list catalog
		return ""
	}
	act := ""
	if a.action == rbac.ActionPull {
		act = "pull"
	} else if a.action == rbac.ActionPush {
		act = "pull,push"
	} else if a.action == rbac.ActionDelete {
		act = "delete"
	} else {
		return ""
	}
	return fmt.Sprintf("repository:%s:%s", a.name, act)
}
