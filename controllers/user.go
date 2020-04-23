package controllers

import (
	"myEshop/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//UserController struct
type UserController struct {
	beego.Controller
}

//ShowIndex 展示index页面
func (this *UserController) ShowIndex() {
	//1.获取session
	username := this.GetSession("username")
	//2.检查session
	if username != nil {
		this.Data["isLogin"] = true
		this.Data["username"] = username
	} else {
		this.Data["isLogin"] = false
	}
	beego.Info(username)
	this.TplName = "index.html"
}

func (this *UserController) ShowRegister() {
	//1.获取session
	username := this.GetSession("username")
	//2.检查session
	if username != nil {
		this.Data["isLogin"] = true
		this.Data["username"] = username
	} else {
		this.Data["isLogin"] = false
	}
	this.Layout = "layout.html"
	this.TplName = "registered.html"
}

//ShowLogin 展示login页面
func (this *UserController) ShowLogin() {
	//1.获取session
	username := this.GetSession("username")
	//2.检查session
	if username != nil {
		this.Data["isLogin"] = true
		this.Data["username"] = username
	} else {
		this.Data["isLogin"] = false
	}
	this.Layout = "layout.html"
	this.TplName = "login.html"
}

func (this *UserController) RegisterHandler() {
	//1.拿到数据
	username := this.GetString("username")
	pwd := this.GetString("password")
	//2.检查数据
	o := orm.NewOrm()
	user := models.User{}
	user.Uname = username
	user.Pwd = pwd

	err := o.Read(&user, "Uname")
	if err != nil {
		_, err := o.Insert(&user)
		if err != nil {
			this.Data["errmsg"] = err
		}
		this.SetSession("username", username)
		this.ShowIndex()
	}
	this.ShowRegister()
	this.Data["errmsg"] = "Username already exist!"
}

func (this *UserController) LoginHandler() {
	//1.拿到数据
	username := this.GetString("username")
	pwd := this.GetString("password")
	//2.查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}
	user.Uname = username
	err := o.Read(&user, "Uname")
	if err != nil {
		beego.Info("Username does not exist")
		this.TplName = "login.html"
		this.Data["errmsg"] = "Username does not exist!"
		return
	} else if user.Pwd != pwd {
		beego.Info("validation fail")
		this.TplName = "login.html"
		this.Data["errmsg"] = "Password not correct!"
		return
	} else {
		//3.设置session
		this.SetSession("username", username)
		//4.登录跳转
		this.ShowIndex()
		return
	}
}

func (this *UserController) ShowPaid() {
	//1.获取session
	username := this.GetSession("username")
	//2.检查session
	if username != nil {
		this.Data["isLogin"] = true
		this.Data["username"] = username
	} else {
		this.Data["isLogin"] = false
	}
	this.Layout = "layout.html"
	this.TplName = "paid.html"
}

//LogoutHandler 删除当前session，退出登录
func (this *UserController) LogoutHandler() {
	this.DelSession("username")
	this.Redirect("/index", 302)
}
