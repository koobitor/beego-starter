package routers

import (
	"github.com/koobitor/beego-starter/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
