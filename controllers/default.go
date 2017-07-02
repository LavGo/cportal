package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "CPortal"
	c.Data["Email"] = "jiangcw@Ctrip.com"
	c.TplName = "index.html"
}
