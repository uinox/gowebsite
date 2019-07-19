package controllers

import (
	"github.com/astaxie/beego"
)

type TechController struct {
	beego.Controller
}

func (this *TechController) Get() {
	this.Data["IsTech"] = true
	this.TplName = "tech.html"
}
