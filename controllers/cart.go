package controllers

import (
	"fmt"
	"myEshop/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CartController struct {
	beego.Controller
}

//定义一个slice存放pid
var pids = make([]string, 0)

func (c *CartController) AddItem() {
	//1.拿到数据存放到slice中
	pid := c.GetString("pid")
	pids = append(pids, pid)
	beego.Info("pids =", pids)
	//2.设置session
	c.SetSession("pid", pids)
	c.TplName = "index.html"
}

func (c *MainController) ShowCheckout() {
	beego.Info(pids)
	o := orm.NewOrm()

	//声明一个Map储存遍历出来的数据
	var map1 = make(map[int]map[string]string)

	//遍历pids切片取得pid对应的商品
	for i := 0; i < len(pids); i++ {
		fmt.Printf("i= %v v= %v\n", i, pids[i])
		num, _ := strconv.Atoi(pids[i])
		product := models.Product{PID: num}
		err := o.Read(&product)
		if err != nil {
			beego.Info("o.Read err=", err)
		}
		beego.Info("pname =", product.Pname)
		map1[i] = make(map[string]string, 3)
		map1[i]["Pid"] = strconv.Itoa(product.PID)
		map1[i]["Pname"] = product.Pname
		map1[i]["Amount"] = strconv.FormatFloat(product.Amount, 'f', 2, 64)
		map1[i]["Img"] = product.Img
	}
	for k1, v1 := range map1 {
		beego.Info("k1= ", k1)
		for k2, v2 := range v1 {
			beego.Info("k2=%v v2=%v", k2, v2)
		}
	}
	beego.Info(map1)
	c.Data["Product"] = map1
	c.TplName = "checkout.html"

	//2.取出数据库中对应session的商品
	//o := orm.NewOrm()
	//利用反射将interface类型的pids转换为slice类型进行遍历
	// if reflect.TypeOf(pids).Kind() == reflect.Slice {
	// 	p := reflect.ValueOf(pids)
	// 	for i := 0; i < p.Len(); i++ {
	// 		fmt.Printf("i= %v v= %v\n", i, p.Index(i))
	// 		product := models.Product{PID: p.Index(i)}
	// 	}
	// 	//c.Data["product"] = product
	// 	//c.getDataByID(pid)
	// 	//3.显示购物车
	// }
}

func FloatToString(input_num float32) string {
	// to convert a float number to a string
	return strconv.FormatFloat(float64(input_num), 'f', 6, 64)
}
