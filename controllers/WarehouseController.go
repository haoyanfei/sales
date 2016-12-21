package controllers

import (
	"fmt"
	admin "web1/libary/controllers"
	"web1/models"
)

//WarehouseController xx
type WarehouseController struct {
	admin.AdminController
}

//InputList xx
func (p *WarehouseController) InputList() {
	r := new(models.Warehouse)
	result := r.GetInput(1)
	p.Data["Result"] = result
	fmt.Println(result)
	p.TplName = "warehouse/input_list.html"
}

//OutList xx
func (p *WarehouseController) OutList() {
	r := new(models.Warehouse)
	result := r.GetInput(2)
	p.Data["Result"] = result
	fmt.Println(result)
	p.TplName = "warehouse/out_list.html"
}

//Add xx
func (p *WarehouseController) AddInput() {
	p.Data["Unit"] = new(models.Store).All()
	p.TplName = "warehouse/add.html"
}

//Add xx
func (p *WarehouseController) ReduceInput() {
	p.Data["Unit"] = new(models.Store).All()
	p.TplName = "warehouse/reduce.html"
}

//PostAdd xx
func (p *WarehouseController) PostAdd() {
	//保存主表
	//uid,types,status,notes,author,created_at
	warehose := make(map[string]string)
	warehose["uid"] = "0"
	warehose["types"] = p.GetString("types")
	warehose["notes"] = p.GetString("warehouse_notes")
	warehose["author"] = p.GetString("warehouse_author")
	storeId, _ := p.GetInt("store_id")
	warehouseId, _ := new(models.Warehouse).Insert(warehose)

	m := make(map[string][]string)
	m["sku_id"] = p.Input()["sku_id"]
	m["quantity"] = p.Input()["quantity"]
	m["notes"] = p.Input()["notes"]
	types, _ := p.GetInt("types")
	w := new(models.IoWarehouse).Insert(m, warehouseId, types, storeId)
	fmt.Println(w)
	if types == 1 {
		p.Ctx.Redirect(302, "/admin/warehouse/inputList")
	} else if types == 2 {
		p.Ctx.Redirect(302, "/admin/warehouse/outList")
	}

}

//GetIoWarehouse xx
func (p *WarehouseController) GetIoWarehouse() {
	warehouseId, _ := p.GetInt("warehouse_id")
	data := new(models.IoWarehouse).GetById(warehouseId)
	fmt.Println(data)
	p.Data["json"] = data
	p.ServeJSON()
}
