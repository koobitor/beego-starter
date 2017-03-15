package routers

import (
	"github.com/astaxie/beego"
	"github.com/koobitor/beego-starter/controllers"
)

func init() {
	beego.Include(&controllers.AuthController{})
	beego.Router("/", &controllers.MainController{})
}
