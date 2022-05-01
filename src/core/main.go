package main

import (
	"github.com/astaxie/beego"
	"github.com/cowk8s/harbor/src/core/middlewares"
	"github.com/cowk8s/harbor/src/lib/config"
	"github.com/cowk8s/harbor/src/lib/log"
	"github.com/cowk8s/harbor/src/server"
)

func main() {

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = config.SessionCookieName

	log.Info("initializing configurations...")

	server.RegisterRoutes()

	beego.RunWithMiddleWares("", middlewares.Middlewares()...)
}
