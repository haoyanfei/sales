package controllers

import (
	//"github.com/astaxie/beego"
	// "reflect"
	"fmt"
	admin "web1/libary/controllers"
	"web1/models"
)

type MainController struct {
	admin.AdminController
}

func (c *MainController) Get() {

	c.Data["Website"] = "beego.me休息休息下"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) NotFound() {

	role := new(models.Role)

	//	status, err := role.Insert("xxxxx")
	id, name := role.GetId(1)
	fmt.Println(id)
	fmt.Println(name)

	c.TplName = "404.tpl"
}
func (c *MainController) BadServer() {
	role := new(models.Role)
	id, _ := role.Insert("zzzzz")
	fmt.Println(id)
	c.TplName = "500.tpl"
}

func (c *MainController) Login() {
	c.TplName = "login.tpl"
}
