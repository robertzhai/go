package controllers

import (
	"fmt"
	"github.com/astaxie/beego"

	"hello/models"
)

type FoodController struct {
	beego.Controller
}

var arc models.Food

func (c *FoodController) Get() {

	data, err := arc.Query()
	fmt.Println(data, err)

	c.Data["Fooddata"] = data
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "food.tpl"

}
