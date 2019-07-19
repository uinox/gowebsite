package controllers

import (
	"beeblog/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginatController struct {
	beego.Controller
}

type ResData struct {
	Status int
	Name   string
}

func (c *LoginatController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginatController) Post() {

	name := c.Input().Get("name")
	password := c.Input().Get("password")

	user, err := models.GetUser(name)
	fmt.Println(user.Password)
	if err != nil {
		res := &ResData{
			Status: 404,
		}
		c.Data["json"] = res
	} else {
		if user.Password == password {
			
			res := &ResData{
				Status: 200,
				Name:   name,
			}
			c.Data["json"] = res

		} else {
			res := &ResData{
				Status: 403,
			}
			c.Data["json"] = res
		}
	}

	/* if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	} */

	// c.Redirect("/", 301)
	c.ServeJSON()
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck1, err := ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck1.Value

	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
}
