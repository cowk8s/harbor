package v2token

import (
	"context"
	"log"

	"github.com/cowk8s/harbor/src/common/security"
	"github.com/cowk8s/harbor/src/pkg/permission/types"
)

// tokenSecutiryCtx is used for check permission of an internal signed token.
// The intention for this quy is only for support CLI push/pull. It should not be used in other scenario without careful review
// Each request should have a different instance of tokenSecurityCtx
type tokenSecurityCtx struct {
	logger    *log.Logger
	name      string
	accessMap map[string]map[types.Action]struct{}
	ctl       project.Controller
}

func (t *tokenSecurityCtx) Name() string {
	return "v2token"
}

func (t *tokenSecurityCtx) IsAuthenticated() bool {
	return len(t.name) > 0
}

func (t *tokenSecurityCtx) GetUsername() string {
	return t.name
}

func (t *tokenSecurityCtx) IsSysAdmin() bool {
	return false
}

func (t *tokenSecurityCtx) IsSolutionUser() bool {
	return false
}

// New creates instance of token security context based on access list and name
func New(ctx context.Context, name string, access []*registry_token.ResourceActions) security.Context {

}
