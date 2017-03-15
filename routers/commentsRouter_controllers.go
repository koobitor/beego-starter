package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/koobitor/beego-starter/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/koobitor/beego-starter/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/koobitor/beego-starter/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/koobitor/beego-starter/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/koobitor/beego-starter/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/koobitor/beego-starter/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Forgot",
			Router: `/forgot`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
