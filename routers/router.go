package routers

import (
	"myEshop/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.UserController{}, "get:ShowIndex")
	beego.Router("/registered", &controllers.UserController{}, "get:ShowRegister;post:RegisterHandler")
	// beego.Router("/article", &controllers.MainController{}, "get:ShowArticle")
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:LoginHandler")
	// beego.Router("/admin", &controllers.MainController{}, "get:ShowAdmin")
	// beego.Router("/edit", &controllers.MainController{}, "get:ShowEdit;post:HandlerEdit")
	// beego.Router("/post", &controllers.MainController{}, "get:ShowPost;post:HandlerPost")
	// beego.Router("/delete", &controllers.MainController{}, "get:HandlerDelete")
	beego.Router("/logout", &controllers.UserController{}, "get:LogoutHandler")

}
