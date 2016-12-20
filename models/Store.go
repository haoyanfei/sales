package models

import (
	// "fmt"
	//	"reflect"

	"fmt"
	_ "web1/libary/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//TableName xxx
func (r *Store) TableName() string {
	return "store"

}

//GetByID 通过id获取信息
func (r *Store) GetByID(id int) (int, string) {
	o := orm.NewOrm()
	var unit Store
	o.Raw("select * from store where id= ? ", id).QueryRow(&unit)
	return unit.Id, unit.Title

}

//Insert xx
func (r *Store) Insert(title string) (int64, error) {
	o := orm.NewOrm()
	p, _ := o.Raw("insert into store(title) value(?)").Prepare()
	res, err := p.Exec(title)
	id, err := res.LastInsertId()
	p.Close()
	return id, err
}

//Update xx
func (r *Store) Update(id int, title string) (int64, error) {
	o := orm.NewOrm()
	p, _ := o.Raw("update store set title= ? where id= ? ").Prepare()
	res, err := p.Exec(title, id)
	p.Close()
	rt, err := res.RowsAffected()
	return rt, err
}

//All xx
func (r *Store) All() []Store {
	o := orm.NewOrm()
	var Store []Store
	o.Raw("select id,title from store").QueryRows(&Store)
	fmt.Println(Store)
	return Store

}
