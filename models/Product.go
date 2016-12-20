package models

import (
	// "fmt"
	//	"reflect"

	_ "web1/libary/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//TableName xxx
func (r *Product) TableName() string {
	return "product"

}

//GetByID 通过id获取信息
func (r *Product) GetByID(id int) Product {
	o := orm.NewOrm()
	var product Product
	o.Raw("select * from product where id= ? ", id).QueryRow(&product)
	return product

}

/**
type Product struct {
	Id            int
	Price         int
	CostPrice     int
	ProductNumber string `orm:"size(100)"`
	Title         string `orm:"size(100)"`
	Subtitle      string `orm:"size(100)"`
	Describute    string `orm:"size(100)"`
	CreatedAt     int
	UnitId        int
}
*/
//Insert xx
func (r *Product) Insert(ProductData *Product) (int64, error) {
	o := orm.NewOrm()
	p, _ := o.Raw("insert into product(title,subtitle,describute,created_at,unit_id) values(?,?,?,?,?)").Prepare()
	res, err := p.Exec(ProductData.Title, ProductData.Subtitle, ProductData.Describute, ProductData.CreatedAt, ProductData.UnitId)
	id, err := res.LastInsertId()
	p.Close()
	return id, err
}

//Update xx
func (r *Product) Update(ProductData *Product) (int64, error) {
	o := orm.NewOrm()
	p, _ := o.Raw("update product set title=?,subtitle=?,describute=?,unit_id=? where id= ? ").Prepare()
	res, err := p.Exec(ProductData.Title, ProductData.Subtitle, ProductData.Describute, ProductData.UnitId, ProductData.Id)
	p.Close()
	rt, err := res.RowsAffected()
	return rt, err

}

//All xx
func (r *Product) All() []Product {
	o := orm.NewOrm()
	var product []Product
	o.Raw("select * from product").QueryRows(&product)

	return product

}
