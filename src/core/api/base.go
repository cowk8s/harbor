package api

import (
	"github.com/cowk8s/harbor/src/common/api"
)

const (
	userSessionKey = "user"
)

type BaseController struct {
	api.BaseAPI
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
