package api

import "github.com/cowk8s/harbor/src/common/api"

// BaseController ...
type BaseController struct {
	api.BaseAPI
	// SecurityCtx is the security context used to authN &authZ
}

// Prepare inits security context and project manager from request
// context
func (b *BaseController) Prepare() {

}
