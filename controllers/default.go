//Package controllers is controllers
//for deal route params
package controllers

import (
	//"github.com/astaxie/beego"
	// "reflect"
	"fmt"
	admin "web1/libary/controllers"
	"web1/models"
)

// MainController main
type MainController struct {
	admin.AdminController
}

//Get is all get mothod
func (c *MainController) Get() {

	c.Data["Website"] = "beego.me休息休息下"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// NotFound 404
func (c *MainController) NotFound() {

	role := new(models.Role)

	//	status, err := role.Insert("xxxxx")
	id, name := role.GetId(1)
	fmt.Println(id)
	fmt.Println(name)

	c.TplName = "404.tpl"
}

//BadServer 502
func (c *MainController) BadServer() {
	role := new(models.Role)
	id, _ := role.Insert("zzzzz")
	fmt.Println(id)
	c.TplName = "500.tpl"
}

//Login frame
func (c *MainController) Login() {
	c.TplName = "login.tpl"
}

//Login frame
func (c *MainController) ToLogin() {
	username := c.GetString("username")
	pwd := c.GetString("password")
	fmt.Println(username)
	fmt.Println("xxxxfffff")

	if username == "im" && pwd == "huanhuan123" {
		c.Ctx.SetCookie("userinfo", "goodluck")
		c.Ctx.Redirect(302, "/")
	} else {
		c.Ctx.Redirect(302, "/login")
	}

}
