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

	beego.Router("/admin/role/add", &controllers.RoleController{}, "*:AddAndEdit")
}
