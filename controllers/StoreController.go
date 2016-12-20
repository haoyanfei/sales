package controllers

import (
	"fmt"
	admin "web1/libary/controllers"
	"web1/models"
)

//StoreController xx
type StoreController struct {
	admin.AdminController
}

//AddAndEdit xx
func (p *StoreController) AddAndEdit() {

	UnitID, _ := p.GetInt("Id")

	Title := p.GetString("Title")
	fmt.Println(Title)
	r := new(models.Store)
	if UnitID > 0 {
		flag, _ := r.Update(UnitID, Title)
		fmt.Println(flag)
	} else {
		id, _ := r.Insert(Title)
		fmt.Println(id)
	}
	p.Ctx.Redirect(302, "/admin/store/list")
}

//Add xx
func (p *StoreController) Add() {

	p.TplName = "productUnit/add.tpl"
}

//Edit xx
func (p *StoreController) Edit() {

}

//List xx
func (p *StoreController) List() {
	r := new(models.Store)
	result := r.All()
	p.Data["Result"] = result
	p.TplName = "product_unit/list.html"
}
