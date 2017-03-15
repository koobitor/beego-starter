package controllers

import (
	"github.com/astaxie/beego"
)

// AuthController struct
type AuthController struct {
	beego.Controller
}

// URLMapping Func
func (c *AuthController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Logout", c.Logout)
	c.Mapping("Forgot", c.Forgot)
}

// Login Page
// @Description Login Page
// @router /login [get]
func (c *AuthController) Login() {
	c.TplName = "auth/login.tpl"
	c.Data["title"] = "Login"
}

// Login Page
// @Description Login Page
// @router /login [post]

// Logout Path
// @Description Logout Path
// @router /logout [get]
func (c *AuthController) Logout() {
	c.TplName = "auth/logout.tpl"
	c.Data["title"] = "Logout"
}

// Forgot Password
// @Description Forgot Password
// @router /forgot [get]
func (c *AuthController) Forgot() {
	c.TplName = "auth/forgot.tpl"
	c.Data["title"] = "Forgot Password"
}

// Forgot Password
// @Description Forgot Password
// @router /forgot [POST]
