package controllers

import (
	"myEshop/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CartController struct {
	beego.Controller
}

//定义一个slice存放pid
var cart = make([]int, 0)

func (c *CartController) AddItem() {
	//1.拿到数据存放到slice中
	pid := c.GetString("pid")
	pidc, _ := strconv.Atoi(pid)
	cart = append(cart, pidc)
	beego.Info("cart =", cart)

	c.TplName = "index.html"
}

func (c *MainController) ShowCheckout() {
	beego.Info(cart)
	o := orm.NewOrm()

	//1.声明一个Map储存遍历出来的数据
	var map1 = make(map[int]map[string]string)

	//取出 cart 中每个商品的数量
	counterMap := make(map[int]int)
	for _, v := range cart {
		if counterMap[v] != 0 {
			counterMap[v]++
		} else {
			counterMap[v] = 1
		}
	}
	beego.Info(counterMap)

	//将取出的数量放入 map1[i]["Quantity"]
	//因为 counterMap 中 index 是string类型无法遍历到map1所以声明一个自增的int类型 i 来使用遍历
	i := 1
	for _, v := range counterMap {

		num := cart[i]
		//3.获取特定商品数据
		product := models.Product{PID: num}
		err := o.Read(&product)
		if err != nil {
			beego.Info("o.Read err=", err)
		}
		beego.Info("pname =", product.Pname)

		map1[i] = make(map[string]string, 4)
		map1[i]["Pid"] = strconv.Itoa(product.PID)
		map1[i]["Pname"] = product.Pname
		map1[i]["Amount"] = strconv.FormatFloat(product.Amount, 'f', 2, 64)
		map1[i]["Img"] = product.Img
		map1[i]["Quantity"] = strconv.Itoa(v)
		i++
	}
	beego.Info(map1[1]["Quantity"], map1[2]["Quantity"])
	beego.Info(map1)

	// //2.遍历 cart 切片取得 pid 对应的商品
	// for i := 0; i < len(cart); i++ {
	// 	fmt.Printf("i= %v v= %v\n", i, cart[i])

	// 	num, _ := strconv.Atoi(cart[i])
	// 	//3.获取特定商品数据
	// 	product := models.Product{PID: num}
	// 	err := o.Read(&product)
	// 	if err != nil {
	// 		beego.Info("o.Read err=", err)
	// 	}
	// 	beego.Info("pname =", product.Pname)
	// 	//4.将获取到的数据写入 map
	// 	map1[i] = make(map[string]string, 4)
	// 	map1[i]["Pid"] = strconv.Itoa(product.PID)
	// 	map1[i]["Pname"] = product.Pname
	// 	map1[i]["Amount"] = strconv.FormatFloat(product.Amount, 'f', 2, 64)
	// 	map1[i]["Img"] = product.Img
	// }

	//5.将 map 传到前端页面
	c.Data["Product"] = map1
	c.TplName = "checkout.html"
}
