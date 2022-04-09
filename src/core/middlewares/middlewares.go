package middlewares

import (
	"regexp"

	"github.com/astaxie/beego"
)

var (
	match = regexp.MustCompile
	numericRegexp = match(`[0-9]+`)
)

func MiddleWares() []beego.MiddleWare {
	return []beego.MiddleWare{
		
	}
}