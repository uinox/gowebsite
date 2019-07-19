package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {

	op := this.Input().Get("op")

	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	}

	this.Data["IsCategory"] = true
	this.TplName = "category.html"

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories

	//
	for _, category := range categories {
		topics, err := models.GetAllTopics(category.Title, "", true)
		if err != nil {
			beego.Error(err.Error)
		}

		if len(topics) > 0 {
			o := orm.NewOrm()

			cate := new(models.Category)

			qs := o.QueryTable("category")
			err = qs.Filter("title", category.Title).One(cate)
			if err != nil {
				beego.Error(err.Error)
			}

			cate.TopicCount = int64(len(topics))
			_, err = o.Update(cate)
		}
	}

}
