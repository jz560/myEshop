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
	// beego.Router("/edit", &controllers.MainController{}, "get:ShowEdit;post:HandlerEdit")
	beego.Router("/addItem", &controllers.CartController{}, "post:AddItem")
	beego.Router("/removeItem", &controllers.MainController{}, "post:RemoveItem")
	// beego.Router("/delete", &controllers.MainController{}, "get:HandlerDelete")
	beego.Router("/logout", &controllers.UserController{}, "get:LogoutHandler")

}
