package models

import (
	// "fmt"
	//	"reflect"

	"fmt"
	"strconv"
	"time"
	_ "web1/libary/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
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
*/
//TableName xxx
func (r *Warehouse) TableName() string {
	return "warehouse"
}

//All xx
func (r *Warehouse) All() []orm.Params {
	o := orm.NewOrm()
	var params []orm.Params
	o.Raw("select * from warehouse as A left join sku as B on A.sku_id=B.id left join product as C on C.id=B.product_id").Values(&params)
	fmt.Println(params)
	return params
}

//GetInput xx
func (r *Warehouse) GetInput(types int) []orm.Params {
	o := orm.NewOrm()
	var params []orm.Params
	o.Raw("select * from warehouse where types=?", types).Values(&params)
	return params
}

//Insert xxx
func (r *Warehouse) Insert(m map[string]string) (int64, error) {
	o := orm.NewOrm()
	p, err := o.Raw("insert into warehouse (uid,types,status,notes,author,created_at)values(?,?,?,?,?,?)").Prepare()
	types, _ := strconv.Atoi(m["types"])
	status, _ := strconv.Atoi(m["status"])
	res, err := p.Exec(0, types, status, m["notes"], m["author"], time.Now().Unix())
	id, err := res.LastInsertId()
	return id, err
}

//TableName xx
func (Io *IoWarehouse) TableName() string {
	return "ioWarehouse"
}

//IoWarehouse Insert
func (Io *IoWarehouse) Insert(m map[string][]string, warehouseId int64, types int, storeId int) bool {
	o := orm.NewOrm()
	p, _ := o.Raw("insert into ioWarehouse(warehouse_id,sku_id,quantity,notes,store_id) values(?,?,?,?,?)").Prepare()
	for i := 0; i < len(m["sku_id"]); i++ {
		skuId, _ := strconv.Atoi(m["sku_id"][i])
		quantity, _ := strconv.Atoi(m["quantity"][i])
		p.Exec(warehouseId, skuId, quantity, m["notes"][i], storeId)
		if types == 1 {
			new(Stock).DealStock(storeId, skuId, +quantity)
		} else if types == 2 {
			new(Stock).DealStock(storeId, skuId, -quantity)
		}
	}
	p.Close()
	return true
}

//Discard
func (Io *IoWarehouse) Discard(warehouseId int) bool {
	o := orm.NewOrm()
	//0.可用 1 废弃
	p, _ := o.Raw("update warehouse set `status` = ? , system_notes = ? where id=? ").Prepare()
	p.Exec(1, "discard", warehouseId)
	var args []IoWarehouse
	o.Raw("select id,store_id,sku_id,quantity from ioWarehouse where warehouse_id=?", warehouseId).QueryRows(&args)
	for _, v := range args {
		new(Stock).DiscardStock(v.StoreId, v.SkuId, -v.Quantity)
	}
	p.Close()
	return true
}

//GetById xx
func (Io *IoWarehouse) GetById(warehouseId int) []orm.Params {
	o := orm.NewOrm()
	// var args1, args2 []orm.Params
	var args1 []orm.Params
	o.Raw("select A.id as ioId,A.warehouse_id,A.notes as io_notes,A.sku_id as sku_id,A.quantity as quantity ,B.name as sku_name,B.product_code,B.product_number ,C.title from ioWarehouse as A left join sku as B on B.id =A.sku_id left join product as C on B.product_id=C.id  where A.warehouse_id=?", warehouseId).Values(&args1)
	return args1
}
