package main

import (
	_ "myEshop/routers"

	_ "myEshop/models"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
