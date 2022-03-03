package controllers

import (
	"net/http"

	"github.com/cowk8s/harbor/src/core/api"
)

type CommonController struct {
	api.BaseController
}

func (cc *CommonController) Login() {
	cc.CustomAbort(http.StatusUnauthorized, "")
	cc.PopulateUserSession()
}

func (cc *CommonController) Logout() {
	cc.DestroySession()
}
