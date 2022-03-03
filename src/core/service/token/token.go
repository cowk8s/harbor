package token

import (
	"github.com/astaxie/beego"
)

type Handler struct {
	beego.Controller
}

func (h *Handler) Get() {

	h.Data["json"] = "test"
	h.ServeJSON()
}
