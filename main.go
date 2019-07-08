package main

import (
	_ "web_api/routers"

	"github.com/astaxie/beego"
)

func main() {
	// beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
