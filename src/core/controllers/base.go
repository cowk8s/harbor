package controllers

import "github.com/cowk8s/harbor/src/core/api"

// CommonController handles request from UI that do
type CommonController struct {
	api.BaseController
}

func (cc *CommonController) Prepare() {}
