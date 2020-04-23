package controllers

import (
	"fmt"
	"myEshop/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
)

type CartController struct {
	beego.Controller
}

//定义一个slice存放pid
var cart = make([]int, 0)

func (c *CartController) AddItem() {
	//拿到数据存放到slice中
	pid := c.GetString("pid")
	pidc, _ := strconv.Atoi(pid)
	cart = append(cart, pidc)
	beego.Info("cart =", cart)

	c.TplName = "index.html"
}

func (c *MainController) ShowCheckout() {
	username := c.GetSession("username")
	if username != nil {
		c.Data["isLogin"] = true
		c.Data["username"] = username
	} else {
		c.Data["isLogin"] = false
	}
	beego.Info(cart)
	c.getCart(cart)
}

func (c *MainController) RemoveItem() {
	pid := c.GetString("pid")
	pidc, _ := strconv.Atoi(pid)
	//从cart切片中删掉指定pid的商品
	cart1 := make([]int, 0, len(cart))
	for _, val := range cart {
		if val != pidc {
			cart1 = append(cart1, val)
		}
	}
	beego.Info(cart)
	beego.Info(cart1)
	c.getCart(cart1)
}

func (c *MainController) getCart(cart []int) {
	beego.Info(cart)
	o := orm.NewOrm()

	//1.声明一个Map储存遍历出来的数据
	var map1 = make(map[int]map[string]string)

	//2.取出 cart 中每个商品的数量
	counterMap := make(map[int]int)
	for _, v := range cart {
		if counterMap[v] != 0 {
			counterMap[v]++
		} else {
			counterMap[v] = 1
		}
	}
	beego.Info(len(counterMap))
	beego.Info(counterMap)
	//3.检查counterMap中是否有值，没有则代表购物车为空
	if len(counterMap) == 0 {
		c.Data["IsEmpty"] = true
		c.TplName = "checkout.html"
	}
	beego.Info(c.Data["IsEmpty"])

	var total float64 //所有商品总价
	for i, v := range counterMap {

		PID := i

		//4.查询获取特定商品数据
		product := models.Product{PID: PID}
		err := o.Read(&product)
		if err != nil {
			beego.Info("o.Read err=", err)
		}
		//5.如果查询返回数据，写入map
		if err != orm.ErrNoRows {
			map1[i] = make(map[string]string, 5)
			map1[i]["Pid"] = strconv.Itoa(product.PID)
			map1[i]["Pname"] = product.Pname
			map1[i]["Amount"] = strconv.FormatFloat(product.Amount, 'f', 2, 64)
			map1[i]["Img"] = product.Img
			map1[i]["Quantity"] = strconv.Itoa(v)
			//取得单个商品的总价
			//先将物品单价product.Amount和数量v相乘
			itemTotal := decimal.NewFromFloat(product.Amount).Mul(decimal.NewFromFloat(float64(v)))
			//得到itemTotal是decimal类型，将其转换为float64类型
			itemTotalF, _ := itemTotal.Float64()
			//再将其转换为string类型，就可以写入map当中了
			map1[i]["ItemTotal"] = strconv.FormatFloat(itemTotalF, 'f', 2, 64)

			beego.Info("pname =", product.Pname)
			beego.Info(map1)
		}

	}
	//6.取得商品总价
	for i, _ := range map1 {
		ItemTotal, _ := strconv.ParseFloat(map1[i]["ItemTotal"], 64)
		total += ItemTotal
		beego.Info(total)
	}
	total, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", total), 64)
	//7.将 map 和商品总价传到前端页面
	c.Data["Total"] = total
	c.Data["Product"] = map1
	c.TplName = "checkout.html"
}
