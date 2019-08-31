package controllers

import (
"beeblog/models"
"path"
"fmt"

"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {

	tid := this.Input().Get("name")

	if len(tid) == 0 {
		topics, err := models.GetAllTopics(this.Input().Get("cate"), "", true)
		if err != nil {
			beego.Error(err.Error)
		}

		this.Data["json"] = topics

		this.ServeJSON()

	} else {
		topic, err := models.GetTopic(tid)

		if err != nil {
			beego.Error(err)
			return
		}
		this.Data["json"] = topic

		this.ServeJSON()
	}

	return
}

func (this *UploadController) Post() {

	//获取附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh != nil {
		//保存附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}
	fmt.Println(attachment)
	this.Data["json"] = "/apis/attachment/"+attachment
	this.ServeJSON()

	return
}

func (this *UploadController) Delete() {
	/* if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	} */

	err := models.DeleteTopic(this.Input().Get("tid"), this.Input().Get("category"))
	if err != nil {
		beego.Error(err)
	}

	this.Data["json"] = "[{status:'successfull'}]"
	this.ServeJSON()

	return
}


