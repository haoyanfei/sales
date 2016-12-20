package models

import (
	// "fmt"
	//	"reflect"

	_ "web1/libary/models"

	_ "github.com/go-sql-driver/mysql"
)

/*
//TableName xxx
func (r *Category) TableName() string {
	return "category"

}

//GetByID 通过id获取信息
func (r *Category) GetByID(id int) (int, string) {
	o := orm.NewOrm()
	var unit ProductUnit
	o.Raw("select * from category where id= ? ", id).QueryRow(&unit)
	return unit.Id, unit.Title

}

//Insert xx
func (r *Category) Insert(title string) (int64, error) {
	o := orm.NewOrm()
	p, _ := o.Raw("insert into category(title) value(?)").Prepare()
	res, err := p.Exec(title)
	id, err := res.LastInsertId()
	p.Close()
	return id, err
}

//Update xx
func (r *Category) Update(id int, title string) (int64, error) {
	o := orm.NewOrm()
	p, _ := o.Raw("update category set title= ? where id= ? ").Prepare()
	res, err := p.Exec(title, id)
	p.Close()
	rt, err := res.RowsAffected()
	return rt, err
}

//All xx
func (r *Category) All() []ProductUnit {
	o := orm.NewOrm()
	var productUnit []ProductUnit
	o.Raw("select * from category").QueryRows(&productUnit)
	fmt.Println(productUnit)
	return productUnit

}
*/
