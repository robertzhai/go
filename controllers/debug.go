package controllers

import (
	"github.com/astaxie/beego"
)

type DebugController struct {
	beego.Controller
}

func (c *DebugController) Get() {
	c.Ctx.WriteString("debug beego")
}
