package controllers

import (
	"beeblog/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GetCategoryAtController struct {
	beego.Controller
}

func (this *GetCategoryAtController) GetCategoryAt() {

	// op := this.Input().Get("op")

	var ss = &models.Tm{}
	json.Unmarshal(this.Ctx.Input.RequestBody, ss)
	op := ss.Op

	switch op {
	case "add":
		// name := this.Input().Get("name")
		name := ss.Name
		fmt.Println(name)
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

	case "del":
		// id := this.Input().Get("id")
		id := ss.Id
		// fmt.Println(id)
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	// this.Data["Categories"] = categories

	this.Data["json"] = categories

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

	this.ServeJSON()
	return
}
