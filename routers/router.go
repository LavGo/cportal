package routers

import (
	"github.com/LavGo/cportal/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/v1/infos.do",&controllers.ApiMainController{})
}
