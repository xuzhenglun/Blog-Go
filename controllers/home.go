package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "index.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	topics, err := models.GetAllTopics(this.Input().Get("cate"), true)
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
}
