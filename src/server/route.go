package server

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/cowk8s/harbor/src/core/controllers"
	"github.com/cowk8s/harbor/src/server/router"
)

func registerRoutes() {
	// API version
	router.NewRoute().Method(http.MethodGet).Path("/api/version").HandlerFunc(GetAPIVersion)

	// Controller API:
	beego.Router("/c/login", &controllers.CommonController{}, "post:Login")
}
