package models

/**
sql:
	create database adminManger default charset utf8 collate utf8-general_ci;


	create table menu (	id int primary key auto_increment,
							title varchar(50),
							parent_id int,
							`level` tinyint,
							created_at int
						);
	create table permission(	id int primary key auto_increment,
								menu_id int,
								mca int,
								is_nav tinyint,
								created_at int
							);
	create table role (	id int primary key auto_increment,
							name varchar(50)
						);
	create table rolePermission( 	id int primary key auto_increment,
									role_id int,
									permission_id int,
									created_at int
								);
	create table userRole (	id int primary key auto_increment,
								uid int,
								role_id int
							);
	create table user ( 	id int primary key auto_increment,
							username varchar(100),
							email varchar(100),
							created_at int
						);


*/

import (
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	Id         int
	Title      string `orm:"size(100)"`
	Parent_id  int
	Created_at int
	Level      int
}
type Permission struct {
	Id         int
	Menu_id    int
	Mca        string `orm:"size(50)"`
	Is_nav     int
	Created_at int
}

type Role struct {
	Id   int
	Name string `orm:"size(100)"`
}
type RolePermission struct {
	Id            int
	Role_id       int
	Permission_id int
	Create_at     int
}
type UserRole struct {
	Id      int
	Uid     int
	Role_id int
}
type User struct {
	Id         int
	Username   string `orm:"size(100)"`
	Email      string `orm:"size(100)"`
	Created_at int
}

func init() {
	orm.RegisterModel(new(User), new(Permission), new(Role), new(RolePermission), new(UserRole))
}
