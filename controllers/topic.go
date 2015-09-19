package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xuzhenglun/Blog-Go/models"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")
	category := this.Input().Get("category")
	tags := this.Input().Get("tags")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, tags, content)
	} else {
		err = models.ModifyTopic(tid, title, category, tags, content)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplNames = "topic.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}

}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplNames = "topic_add.html"
}

func (this *TopicController) View() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplNames = "topic_view.html"

	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	tid := this.Ctx.Input.Param("0")
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
	this.Data["Tag"] = strings.Split(topic.Tag, ",")

	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
}

func (this *TopicController) Modify() {
	this.TplNames = "topic_modify.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true

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

	tid := this.Ctx.Input.Param("0")

	err := models.DeleteTopic(tid)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteReplies(tid)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}
