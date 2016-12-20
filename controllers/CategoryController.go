package controllers

/*
//CategoryController xx
type CategoryController struct {
	admin.AdminController
}

//AddAndEdit xx
func (p *CategoryController) AddAndEdit() {

	UnitID, _ := p.GetInt("Id")

	Title := p.GetString("Title")

	r := new(models.ProductUnit)
	if UnitID > 0 {
		flag, _ := r.Update(UnitID, Title)
		fmt.Println(flag)
	} else {
		id, _ := r.Insert(Title)
		fmt.Println(id)
	}
	// p.Ctx.Redirect(302, "/admin/unit/list")
}

//Add xx
func (p *CategoryController) Add() {

	// p.TplName = "productUnit/add.tpl"
}

//Edit xx
func (p *CategoryController) Edit() {

}


//List xx
func (p *CategoryController) List() {
	// r := new(models.ProductUnit)
	result := r.All()
	p.Data["Result"] = result
	p.TplName = "category/list.html"
}

*/
