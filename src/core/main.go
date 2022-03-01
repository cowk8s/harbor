package main

import (
	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = config
}
