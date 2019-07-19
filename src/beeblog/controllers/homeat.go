package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {

	topics, err := models.GetAllTopics(this.Input().Get("cate"), this.Input().Get("label"), true)
	if err != nil {
		beego.Error(err.Error)
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	result := &models.HomeData{
		Topics:     topics,
		Categories: categories,
	}

	this.Data["json"] = result

	this.ServeJSON()

	return

}
