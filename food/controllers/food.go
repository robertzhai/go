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

	data, err := arc.Query(1)
	fmt.Println(data, err)

	c.Data["Fooddata"] = data
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "food.tpl"

}

func (c *FoodController) Ajax() {

    page, _ := c.GetInt("page")
	fmt.Println("ajax................")
	fmt.Println(page)
	data, err := arc.Query(page)
	fmt.Println(data, err)


    c.Data["Fooddata"] = data
	
	c.TplNames = "itemv0.tpl"
	//mystruct := data
 //   c.Data["json"] = &mystruct
 //   c.ServeJson()

}


