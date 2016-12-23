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
		p.Exec(SkuData["name"][i], productId, price*100, cosePrice*100, SkuData["product_number"][i], SkuData["product_code"][i])
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
				p.Exec(SkuData["name"][i], price*100, cosePrice*100, SkuData["product_number"][i], SkuData["product_code"][i], id)
			} else {
				p, _ := o.Raw("insert into sku(name,product_id,price,cost_price,product_number,product_code) values(?,?,?,?,?,?)").Prepare()
				p.Exec(SkuData["name"][i], productId, price*100, cosePrice*100, SkuData["product_number"][i], SkuData["product_code"][i])
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

func (r *Sku) GetStock() []orm.Params {
	o := orm.NewOrm()
	var args []orm.Params
	o.Raw("select A.store_id , B.id as sku_id,B.name as sku_name,C.title as sku_title ,B.product_code as product_code,A.stock as useful from stock as A left join sku as B on A.sku_id=B.id left join product as C on B.product_id=C.id").Values(&args)
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

//DealStock
func (s *Stock) DealStock(store_id int, sku_id int, quantity int) bool {
	//store_id sku_id quantity
	o := orm.NewOrm()
	var stock []Stock
	num, _ := o.Raw("select * from stock where store_id =? AND sku_id=?", store_id, sku_id).QueryRows(&stock)
	var p orm.RawPreparer
	if num > 0 {
		p, _ = o.Raw("update stock set stock = stock + ? where store_id =? AND sku_id=? ").Prepare()
	} else {
		p, _ = o.Raw("insert into stock (stock,store_id,sku_id)values(?,?,?)").Prepare()
	}
	res, _ := p.Exec(quantity, store_id, sku_id)
	p.Close()
	fmt.Println(res)
	return true
}

//DealStock
func (s *Stock) DiscardStock(storeId int, skuId int, quantity int) bool {
	//store_id sku_id quantity
	o := orm.NewOrm()
	p, _ := o.Raw("update stock set stock = stock + ? where store_id =? AND sku_id=? ").Prepare()
	res, _ := p.Exec(quantity, storeId, skuId)
	fmt.Println(res)
	p.Close()
	return true
}
