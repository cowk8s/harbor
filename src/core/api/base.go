package api

import (
	"github.com/cowk8s/harbor/src/common/api"
)

const (
	yamlFileContentType = "application/x-yaml"
	userSessionKey      = "user"
)

// BaseController ...
type BaseController struct {
	api.BaseAPI
	// SecurityCtx is the security context used to authN &authZ
	SecurityCtx
}

func (b *BaseController) Prepare() {

}

func (b *BaseController) RequireAuthenticated() bool {
	return false
}

func (b *BaseController) PopulateUserSession() {
	b.SessionRegenerateID()
	b.SetSession(userSessionKey, "hi")
}
