package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(this.Input().Get("cate"), this.Input().Get("label"), true)
	if err != nil {
		beego.Error(err.Error)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	this.Data["Categories"] = categories

	/* 	this.Data["Website"] = "beego.me"
	   	this.Data["Email"] = "astaxie@gmail.com"

	   	this.Data["TrueCond"] = true
	   	this.Data["FalseCond"] = false

	   	type u struct {
	   		Name string
	   		Age  int
	   		Sex  string
	   	}

	   	user := &u{
	   		Name: "Joe",
	   		Age:  20,
	   		Sex:  "Male",
	   	}

	   	this.Data["User"] = user

	   	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	   	this.Data["Nums"] = nums

	   	this.Data["TplVar"] = "hey guys"

	   	this.Data["html"] = "<div>hello beego</div>"

	   	this.Data["Pipe"] = "<div>hello beego</div>" */
}
