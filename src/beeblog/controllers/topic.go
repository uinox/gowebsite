package controllers

import (
	"beeblog/models"
	"fmt"
	"path"
	"strings"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err.Error)
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

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
	fmt.Println(uid)

	// var err error
	if len(tid) == 0 {
		err = models.AddTopic(uid, title, category, label, content, author, attachment)
	} else {
		err = models.ModifyTopic(tid, title, category, label, content, attachment)
	}
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"

	/*
		reqUrl := this.Ctx.Request.RequestURI
		i := strings.LastIndex(reqUrl, "/")
		tid := reqUrl[i+1:]
		topic, err := models.GetTopic(tid)
	*/
	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(topic.Labels, " ")
	// this.Data["Tid"] = this.Ctx.Input.Param("0")

	replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	fmt.Println(len(replies))
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)

	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid

}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"), this.Input().Get("category"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/", 302)
}
