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

	var Person models.User

	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal([]byte(data),&Person)
	if err != nil {
		fmt.Println(err)
	}

	name := Person.Name
	password := Person.Password
	email := Person.Email


	state := models.AddUser(name, password, email)
	fmt.Println(state)
	if state == 301 {
		this.Data["json"] = &JSON{"301","用户已存在"}
	}
	if state == 304 {
		this.Data["json"] = &JSON{"304","添加用户失败"}
	}
	if state == 200 {
		this.Data["json"] = &JSON{"200", "添加成功"}
	}
	this.ServeJSON()
	return
}
