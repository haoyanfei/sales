package models

import (
	"fmt"
	//	"reflect"
	_ "web1/libary/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func (r *Role) GetId(id int) (int, string) {
	o := orm.NewOrm()
	var role Role
	o.Raw("select * from role where id= ? ", id).QueryRow(&role)
	return role.Id, role.Name

}

func (r *Role) Insert(name string) (int64, error) {
	o := orm.NewOrm()
	p, _ := o.Raw("insert into role(name) value(?)").Prepare()
	res, err := p.Exec(name)
	id, err := res.LastInsertId()
	p.Close()
	return id, err
}
func (r *Role) Update(id int, name string) (int64, error) {
	o := orm.NewOrm()
	p, _ := o.Raw("update role set name= ? where id= ? ").Prepare()
	res, err := p.Exec(name, id)
	p.Close()
	rt, err := res.RowsAffected()
	return rt, err
}
