package server

import (
	"github.com/astaxie/beego"

	"github.com/cowk8s/harbor/src/core/controllers"
	"github.com/cowk8s/harbor/src/core/service/token"
)

func registerRoutes() {
	// Controller API:
	beego.Router("/c/login", &controllers.CommonController{}, "get:Login")

	beego.Router("/service/token", &token.Handler{})
}
