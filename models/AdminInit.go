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
	create table ProductUnit (id int primary key auto_increment,
							title varchar(100),
						);

*/

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

//User xx
type User struct {
	ID        int
	Username  string `orm:"size(100)"`
	Email     string `orm:"size(100)"`
	CreatedAt int
}

//ProductType xx
type ProductType struct {
	ID    int
	Title string `orm:"size(100)"`
}

//ProductUnit 计量单位
type Store struct {
	Id    int
	Title string `orm:"size(100)"`
}

//Category 分类
type Category struct {
	Id        int
	Parent_id int
	Level     int
	Sort      int
	Title     string `orm:"size(100)"`
}

//Product 商品
type Product struct {
	Id            int
	Price         int
	CostPrice     int
	ProductNumber string `orm:"size(100)"`
	Title         string `orm:"size(100)"`
	Subtitle      string `orm:"size(100)"`
	Describute    string `orm:"size(100)"`
	CreatedAt     int64
	UnitId        int
}

//Sku 规格
type Sku struct {
	Id            int
	Name          string `orm:"size(100)"`
	ProductId     int
	Price         int
	CostPrice     int
	ProductNumber string `orm:"size(100)"`
	ProductCode   string `orm:"size(100)"`
	CreatedAt     int
	Stock         int
	Lock          int
}

type Stock struct {
	Id       int
	SkuId    int
	Stock    int
	StoreId  int
	CreateAt int
}

/**
CREATE TABLE `golang`.`Stock` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` INT NULL,
  `stock` INT NULL,
  `lock` INT NULL,
  PRIMARY KEY (`id`));
*/

/**
CREATE TABLE `golang`.`warehouse` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` INT UNSIGNED NULL,
  `quantity` INT UNSIGNED NULL,
  `type` INT UNSIGNED NULL,
  `status` INT UNSIGNED NULL,
  PRIMARY KEY (`id`));
*/

type IoWarehouse struct {
	Id          int
	WarehouseId int
	Uid         int
	SkuId       int
	Quantity    int
	Notes       string `orm:"size(100)"`
}

//Warehouse 入库
type Warehouse struct {
	Id        int
	ProductId int
	SkuId     int
	Quantity  int
	Types     int    //1.入库 2.出库
	Status    int    // 0.待审核.1审核.2. 驳回
	Notes     string `orm:"size(100)"`
	Author    string `orm:"size(100)"`
	CreatedAt int
}
