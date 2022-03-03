package main

import (
	"github.com/astaxie/beego"
	"github.com/cowk8s/harbor/src/lib/log"
	"github.com/cowk8s/harbor/src/server"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "hi"
	log.Info("initializing configurations...")

	server.RegisterRoutes()

	beego.Run()
}
