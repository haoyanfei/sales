package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// default addr for network '172.16.216.185' unknown
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/golang?charset=utf8")
	orm.Debug = true
}
