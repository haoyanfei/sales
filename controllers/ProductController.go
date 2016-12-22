package controllers

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	admin "web1/libary/controllers"
	"web1/models"
)

//ProductController xx
type ProductController struct {
	admin.AdminController
}

//AddAndEdit xx
func (p *ProductController) AddAndEdit() {

	fmt.Println(p.Input())

	ProductData := new(models.Product)

	ProductId, _ := p.GetInt("product_id")
	ProductData.Title = p.GetString("title")
	ProductData.Describute = p.GetString("describute")
	ProductData.Subtitle = p.GetString("subtitle")
	ProductData.UnitId, _ = p.GetInt("unit_id")
	ProductData.CreatedAt = time.Now().Unix()
	m := make(map[string][]string, 0)

	m["price"] = p.Input()["price"]
	m["name"] = p.Input()["name"]
	m["cost_price"] = p.Input()["cost_price"]
	m["product_number"] = p.Input()["product_number"]
	m["product_code"] = p.Input()["product_code"]

	r := new(models.Product)
	sku := new(models.Sku)
	if ProductId > 0 {
		ProductData.Id = ProductId
		flag, _ := r.Update(ProductData)
		m["id"] = p.Input()["sku_id"]
		bool := sku.Update(m, ProductId)
		fmt.Println(flag)
		fmt.Println(bool)
	} else {
		id, _ := r.Insert(ProductData)
		// fmt.Println(id)
		bool := sku.Insert(m, id)
		fmt.Println(bool)
	}
	p.Ctx.Redirect(302, "/admin/product/list")
}

//List xx
func (p *ProductController) List() {
	r := new(models.Product)
	var w bytes.Buffer
	productCode := p.GetString("product_code")
	w.WriteString("1=1 ")
	if productCode != "" {
		w.WriteString(" AND product_code like \"%")
		w.WriteString(productCode)
		w.WriteString("%\" ")
	}
	productNumber := p.GetString("product_number")
	if productNumber != "" {
		w.WriteString("AND product_number like \"%")
		w.WriteString(productNumber)
		w.WriteString("%\" ")
	}
	title := p.GetString("title")
	if title != "" {
		w.WriteString("AND title like \"%")
		w.WriteString(title)
		w.WriteString("%\" ")
	}

	where := w.String()

	result := r.AllByWhere(where)
	fmt.Println(result)
	p.Data["Result"] = result
	p.TplName = "product/list.html"
}

//Add xx
func (p *ProductController) Add() {

	p.TplName = "product/add.html"
}

//Edit xxx
func (p *ProductController) Edit() {
	id, _ := p.GetInt("id")
	p.Data["Info"] = new(models.Product).GetByID(id)
	p.Data["Sku"] = new(models.Sku).GetByProductId(id)
	p.TplName = "product/edit.html"
}

//Stock xxx
func (p *ProductController) Stock() {
	p.Data["Info"] = new(models.Sku).GetStock()
	fmt.Println(p.Data["Info"])
	p.Data["store"] = new(models.Store).FormatAll()
	fmt.Println("---------------", p.Data["store"])
	p.TplName = "product/stock.html"
}

//ProductForSelected xx
func (p *ProductController) ProductForSelected() {
	sku := new(models.Sku).GetAllSku()
	fmt.Println(sku)

	mp := make([]map[string]interface{}, len(sku))

	var skuName, skuTitle, productCode string

	for k, v := range sku {
		if v["sku_name"] == nil {
			skuName = ""
		} else {
			skuName = v["sku_name"].(string)
		}
		if v["sku_title"] == nil {
			skuTitle = ""
		} else {
			skuTitle = v["sku_title"].(string)
		}
		if v["product_code"] == nil {
			productCode = ""
		} else {
			productCode = v["product_code"].(string)
		}
		mp[k] = make(map[string]interface{}, len(v))
		mp[k]["id"], _ = strconv.Atoi(v["sku_id"].(string))
		mp[k]["text"] = skuTitle + "[" + skuName + "]" + productCode
	}
	fmt.Println(mp)
	p.Data["json"] = mp

	p.ServeJSON()
}
