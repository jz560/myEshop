package routers

import (
	"myEshop/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.UserController{}, "get:ShowIndex")
	beego.Router("/registered", &controllers.UserController{}, "get:ShowRegister;post:RegisterHandler")
	beego.Router("/checkout", &controllers.MainController{}, "get:ShowCheckout")
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:LoginHandler")
	beego.Router("/addItem", &controllers.CartController{}, "post:AddItem")
	beego.Router("/removeItem", &controllers.MainController{}, "post:RemoveItem")
	beego.Router("/logout", &controllers.UserController{}, "get:LogoutHandler")
	beego.Router("/paid", &controllers.UserController{}, "get:ShowPaid")
}
