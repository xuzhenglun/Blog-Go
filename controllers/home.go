package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xuzhenglun/Blog-Go/models"
)

type MainController struct {
	baseController
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "index.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	topics, err := models.GetAllTopics(this.Input().Get("cate"), this.Input().Get("tag"), true)

	for _, topic := range topics {
		if len(topic.Content) > 200 {
			topic.Content = string([]byte(topic.Content)[:200])
		}
	}

	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}

	Categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = Categories

	this.locale() //locale
}
