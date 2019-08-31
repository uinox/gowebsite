package controllers

import (
	"beeblog/models"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
type JSON struct {
    Code string
    Msg  string
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

	//name := this.Input().Get("name")
	//password := this.Input().Get("password")
	//email := this.Input().Get("email")
	var Person models.User

	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal([]byte(data),&Person)
	if err != nil {
		fmt.Println(err)
	}

	name := Person.Name
	password := Person.Password
	email := Person.Email

	fmt.Println("------")
	fmt.Println(name)
	fmt.Println(password)
	fmt.Println(email)

	err = models.AddUser(name, password, email)
	if err != nil {
		beego.Error(err)

		return
	}
	res := &JSON{"200", "获取成功"}
	this.Data["json"] = res
	this.ServeJSON()
	return
}
