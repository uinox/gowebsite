package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type ReplyatController struct {
	beego.Controller
}

func (this *ReplyatController) Add() {
	tid := this.Input().Get("tid")
	err := models.AddReply(tid, this.Input().Get("nickname"), this.Input().Get("content"))

	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = tid
	this.ServeJSON()
	return
}

func (this *ReplyatController) Delete() {
	// if !checkAccount(this.Ctx) {
	// 	return
	// }

	// tid := this.Input().Get("tid")
	rid := this.Input().Get("rid")
	err := models.DeleteReply(rid)

	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = "[{status:'successfull'}]"
	this.ServeJSON()

	return
}
