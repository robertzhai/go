package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/debug", &controllers.DebugController{})
	beego.Router("/food", &controllers.FoodController{})
	beego.Router("/food/ajax", &controllers.FoodController{}, "Get:Ajax")

}
