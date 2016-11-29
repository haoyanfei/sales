package models

import(
  _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"

)
func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
// default addr for network '172.16.216.185' unknown
    orm.RegisterDataBase("default", "mysql", "root:123456@tcp(172.16.216.185:3306)/golang?charset=utf8")
}
