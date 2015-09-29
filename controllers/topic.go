package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/xuzhenglun/Blog-Go/models"
	"io"
	"os"
	"path"
	"strings"
)

type TopicController struct {
	baseController
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

	var attachment string
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	if fh != nil {
		attachment = fh.Filename
		beego.Info(attachment)
		fpath := path.Join("attachment", attachment+".tmp")
		err = this.SaveToFile("attachment", fpath)

		file, err := os.Open(fpath)
		if err != nil {
			attachment = ""
			beego.Error(err)
		} else {
			defer file.Close()
			md5h := md5.New()
			io.Copy(md5h, file)
			if tid == "" {
				tid = "1"
			}
			attachment = fmt.Sprintf("%X", md5h.Sum([]byte(""))[:2]) + "_" + tid + "_" + attachment
			err = os.Rename(fpath, path.Join("attachment", attachment))
			if err != nil {
				beego.Error(err)
			}
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, category, tags, content, attachment)
	} else {
		err = models.ModifyTopic(tid, title, category, tags, content, attachment)
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

	this.locale()
}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplNames = "topic_add.html"
	this.locale()
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
	this.Data["Tid"] = tid
	this.Data["Tag"] = strings.Split(topic.Tag, ",")
	topic.Content = string(blackfriday.MarkdownCommon([]byte(topic.Content)))
	this.Data["Topic"] = topic

	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}

	for _, reply := range replies {
		unsafe := blackfriday.MarkdownCommon([]byte(reply.Content))
		reply.Content = string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))
	}

	this.Data["Replies"] = replies
	this.locale()
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
	this.locale()
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
