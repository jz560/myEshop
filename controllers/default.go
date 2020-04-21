package controllers

import (
	"myEshop/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

//getATable gets all data from product table
func (c *MainController) getPTable(link string) {
	o := orm.NewOrm()
	var product []models.Product
	_, err := o.QueryTable("Product").All(&product)
	if err != nil {
		beego.Info("query table err=", err)
		return
	}
	beego.Info(product)
	c.Data["product"] = product
	c.TplName = link
}

//getATable gets specific id's data from product table
func (c *MainController) getDataByID(PID int) {
	o := orm.NewOrm()
	product := models.Product{PID: PID}
	err := o.Read(&product)
	if err != nil {
		beego.Info("o.Read err=", err)
		return
	}
	c.Data["product"] = product
}
