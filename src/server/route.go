package server

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/cowk8s/harbor/src/common"
	"github.com/cowk8s/harbor/src/core/controllers"
	"github.com/cowk8s/harbor/src/core/service/token"
	"github.com/cowk8s/harbor/src/server/router"
)

func ignoreNotification(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func registerRoutes() {
	// API version
	router.NewRoute().Method(http.MethodGet).Path("/api/version").HandlerFunc(GetAPIVersion)

	// Controller API:
	beego.Router("/c/login", &controllers.CommonController{}, "post:Login")
	beego.Router("/c/log_out", &controllers.CommonController{}, "get:LogOut")
	beego.Router("/c/userExists", &controllers.CommonController{}, "post:UserExists")
	beego.Router(common.OIDCLoginPath, &controllers.OIDCController{}, "get:RedirectLogin")
	beego.Router("/c/oidc/onboard", &controllers.OIDCController{}, "post:Onboard")
	beego.Router(common.OIDCCallbackPath, &controllers.OIDCController{}, "get:Callback")
	beego.Router(common.AuthProxyRediretPath, &controllers.AuthProxyController{}, "get:HandleRedirect")

	beego.Router("/service/token", &token.Handler{})
}
