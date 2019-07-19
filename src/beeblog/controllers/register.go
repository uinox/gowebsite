package controllers

import (
	"beeblog/models"
	"fmt"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
	c.TplName = "login.html"
}

func (this *RegisterController) Post() {

	name := this.Input().Get("name")
	password := this.Input().Get("password")
	email := this.Input().Get("email")
	fmt.Println("------")
	fmt.Println(name)
	fmt.Println(password)
	fmt.Println(email)

	err := models.AddUser(name, password, email)
	if err != nil {
		beego.Error(err)

		return
	}

	this.Data["json"] = "[{'status':'successfull'}]"
	this.ServeJSON()
	return
}
