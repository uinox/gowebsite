package controllers

import (
	"beeblog/models"
	"path"

	"github.com/astaxie/beego"
)

type TopicAtController struct {
	beego.Controller
}
type JSONS struct {
    //必须的大写开头
    Code string
    Msg  string
}

func (this *TopicAtController) Get() {

	tid := this.Input().Get("tid")

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

func (this *TopicAtController) Post() {
	//解析表单
	title := this.Input().Get("title")
	tid := this.Input().Get("tid")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	label := this.Input().Get("label")
	author := this.Input().Get("author")

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

	//获取用户
	user := &models.User{}
	user, err = models.GetUser(author)
	if err != nil {
		beego.Error(err)
	}
	uid := user.Id

	// var err error
	if len(tid) == 0 {
		err = models.AddTopic(uid, title, category, label, content, author, attachment)
	} else {
		err = models.ModifyTopic(tid, title, category, label, content, attachment)
	}
	if err != nil {
		beego.Error(err)
	}
	data := &JSONS{"200","获取成功"}
        this.Data["json"] = data
	this.ServeJSON()

	return
}

func (this *TopicAtController) Delete() {
	/* if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	} */

	err := models.DeleteTopic(this.Input().Get("tid"), this.Input().Get("category"))
	if err != nil {
		beego.Error(err)
	}	
	data := &JSONS{"200","获取成功"}
	this.Data["json"] = data
	this.ServeJSON()

	return
}

func (this *TopicAtController) View() {

	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}

	replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}

	result := &models.TopicView{
		Topic:   topic,
		Replies: replies,
	}

	this.Data["json"] = result
	this.ServeJSON()

	return
}
