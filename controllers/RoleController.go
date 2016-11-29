package controllers

import (
	"fmt"
	admin "web1/libary/controllers"
	"web1/models"
)

type RoleController struct {
	admin.AdminController
}

func (this *RoleController) AddAndEdit() {

	Rid, _ := this.GetInt("Id")
	Name := this.GetString("Name")
	r := new(models.Role)
	if Rid > 0 {
		flag, _ := r.Update(Rid, Name)
		fmt.Println(flag)
	} else {
		id, _ := r.Insert(Name)
		fmt.Println(id)
	}
	this.TplName = "role/add.tpl"
}
