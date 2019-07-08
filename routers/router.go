package routers

import (
	"web_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/parseForm", &controllers.MainController{}, "post:MyParseForm")
	// beego.Router("/parseJson", &controllers.MainController{}, "post:ParseJson")
	ns1 := beego.NewNamespace("/default",
		beego.NSRouter("/parseForm", &controllers.MainController{}, "post:MyParseForm"),
		beego.NSRouter("/parseJson", &controllers.MainController{}, "post:ParseJson"),
	)
	ns2 := beego.NewNamespace("/user",
		beego.NSRouter("/list", &controllers.UserController{}, "get:List"),
	)
	ns3 := beego.NewNamespace("/order",
		beego.NSRouter("/list", &controllers.OrderController{}, "get:List"),
	)
	beego.AddNamespace(ns1)
	beego.AddNamespace(ns2)
	beego.AddNamespace(ns3)
}
