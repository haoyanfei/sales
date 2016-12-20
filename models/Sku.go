package models

import (
	// "fmt"
	//	"reflect"

	"fmt"
	_ "web1/libary/models"

	"strconv"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//TableName xxx
func (r *Sku) TableName() string {
	return "sku"

}

//GetByID 通过id获取信息
func (r *Sku) GetByID(id int) Sku {
	o := orm.NewOrm()
	var sku Sku
	o.Raw("select * from sku where id= ? ", id).QueryRow(&sku)
	return sku

}

/**
type Sku struct {
	Id            int
	Price         int
	CostPrice     int
	SkuNumber string `orm:"size(100)"`
	Title         string `orm:"size(100)"`
	Subtitle      string `orm:"size(100)"`
	Describute    string `orm:"size(100)"`
	CreatedAt     int
	UnitId        int
}
*/
//Insert xx
func (r *Sku) Insert(SkuData map[string][]string, productId int64) bool {
	o := orm.NewOrm()
	p, _ := o.Raw("insert into sku(name,product_id,price,cost_price,product_number,product_code) values(?,?,?,?,?,?)").Prepare()
	fmt.Println(SkuData)
	fmt.Println(len(SkuData["name"]))
	for i := 0; i < len(SkuData["name"]); i++ {
		price, _ := strconv.Atoi(SkuData["price"][i])
		cosePrice, _ := strconv.Atoi(SkuData["cost_price"][i])
		p.Exec(SkuData["name"][i], productId, price, cosePrice, SkuData["product_number"][i], SkuData["product_code"][i])
	}
	p.Close()
	return true
}

//Update xx
func (r *Sku) Update(SkuData map[string][]string, productId int) bool {
	o := orm.NewOrm()
	p, _ := o.Raw("update sku set name=?, price= ? ,cost_price=?,product_number=?,product_code=? where id= ? ").Prepare()
	for i := 0; i < len(SkuData["name"]); i++ {
		price, _ := strconv.Atoi(SkuData["price"][i])
		cosePrice, _ := strconv.Atoi(SkuData["cost_price"][i])
		if ids, ok := SkuData["id"]; ok {
			id, _ := strconv.Atoi(ids[i])
			if id > 0 {
				p.Exec(SkuData["name"][i], price, cosePrice, SkuData["product_number"][i], SkuData["product_code"][i], id)
			} else {
				p, _ := o.Raw("insert into sku(name,product_id,price,cost_price,product_number,product_code) values(?,?,?,?,?,?)").Prepare()
				p.Exec(SkuData["name"][i], productId, price, cosePrice, SkuData["product_number"][i], SkuData["product_code"][i])
			}
		}
	}
	p.Close()
	return true

}

//All xx
func (r *Sku) All() []Sku {
	o := orm.NewOrm()
	var sku []Sku
	o.Raw("select * from sku").QueryRows(&sku)

	return sku

}

//GetByProductId xxx
func (r *Sku) GetByProductId(productId int) []Sku {
	o := orm.NewOrm()
	var sku []Sku
	o.Raw("select * from sku where product_id=?", productId).QueryRows(&sku)
	return sku
}

//GetAllSku xxx
func (r *Sku) GetAllSku() []orm.Params {
	o := orm.NewOrm()
	var args []orm.Params
	o.Raw("select A.id as sku_id,A.name as sku_name,B.title as sku_title ,A.product_code as product_code , stock - `lock` as useful  from sku as A left join product as B on A.product_id=B.id").Values(&args)
	return args
}

//GetSkuByWhere xxx
func (r *Sku) GetSkuByWhere(where string) []orm.Params {
	o := orm.NewOrm()
	fmt.Println(where)
	var args []orm.Params
	o.Raw("select A.id as sku_id,A.name as sku_name,B.title as sku_title ,A.product_code as product_code from sku as A left join product as B on A.product_id=B.id").Values(&args)
	return args
}

//AddStock xx
func (r *Sku) AddStock(skuId int, quantity int) bool {
	o := orm.NewOrm()
	fmt.Println(skuId)
	p, _ := o.Raw("update sku set stock=stock+? where id= ? ").Prepare()
	p.Exec(quantity, skuId)
	p.Close()
	return true
}

//ReduceStock xx
func (r *Sku) ReduceStock(skuId int, quantity int) bool {
	o := orm.NewOrm()
	fmt.Println(skuId)
	p, _ := o.Raw("update sku set stock=stock-? where id= ? ").Prepare()
	p.Exec(quantity, skuId)
	p.Close()
	return true
}
