package models

import (
	"fmt"
	_ "web1/libary/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func (u *User) GetId(id int) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var ut []*User
	qs.Filter("id", 1).All(&ut)
	for _, item := range ut {
		fmt.Println(item.Username)
	}
	// var maps []orm.Params
	// num , _ :=o.Raw("select * from user").Values(&maps)
	// for index,item := range maps{
	//     fmt.Println(index,"----",item)
	// }
	// fmt.Println(num)

}
