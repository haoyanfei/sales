package routers

import (
	"web1/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/404", &controllers.MainController{}, "get:NotFound")
	beego.Router("/500", &controllers.MainController{}, "get:BadServer")

	beego.Router("/login", &controllers.MainController{}, "get:Login")

	beego.Router("/admin/store/addAndEdit", &controllers.StoreController{}, "post:AddAndEdit")

	beego.Router("/admin/store/list", &controllers.StoreController{}, "get:List")

	// beego.Router("/admin/category/list", &controllers.CategoryController{}, "get:List")

	beego.Router("/admin/product/list", &controllers.ProductController{}, "get:List")
	beego.Router("/admin/product/add", &controllers.ProductController{}, "get:Add")
	beego.Router("/admin/product/edit", &controllers.ProductController{}, "get:Edit")
	beego.Router("/admin/product/addAndEdit", &controllers.ProductController{}, "post:AddAndEdit")
	beego.Router("/admin/product/stock", &controllers.ProductController{}, "get:Stock")
	beego.Router("/admin/product/productForSelected", &controllers.ProductController{}, "get:ProductForSelected")

	beego.Router("/admin/warehouse/inputList", &controllers.WarehouseController{}, "get:InputList")
	beego.Router("/admin/warehouse/outList", &controllers.WarehouseController{}, "get:OutList")

	///admin/warehouse/AddInput
	beego.Router("/admin/warehouse/addInput", &controllers.WarehouseController{}, "get:AddInput")
	beego.Router("/admin/warehouse/reduceInput", &controllers.WarehouseController{}, "get:ReduceInput")
	beego.Router("/admin/warehouse/postAdd", &controllers.WarehouseController{}, "post:PostAdd")

	beego.Router("/admin/warehouse/getIoWarehouse", &controllers.WarehouseController{}, "get:GetIoWarehouse")

}
