package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv"
)

type ReplyatController struct {
	beego.Controller
}

func (this *ReplyatController) Add() {
	reply := &models.Comment{}
	json.Unmarshal(this.Ctx.Input.RequestBody,reply)

	tid := strconv.FormatInt(reply.Tid,10)
	nickname := reply.Name
	content := reply.Content

	//tid := this.Input().Get("tid")
	err := models.AddReply(tid, nickname, content)

	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = replies
	this.ServeJSON()
	return
}

func (this *ReplyatController) Delete() {
	// if !checkAccount(this.Ctx) {
	// 	return
	// }

	tid := this.Input().Get("tid")
	rid := this.Input().Get("rid")
	err := models.DeleteReply(rid)

	if err != nil {
		beego.Error(err)
	}
	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = replies
	this.ServeJSON()

	return
}
