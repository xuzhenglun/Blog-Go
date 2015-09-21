package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xuzhenglun/Blog-Go/models"
)

type CategoryController struct {
	baseController
}

func (this *CategoryController) Get() {
	op := this.Input().Get("op")
	switch op {
	case "add":
		if !checkAccount(this.Ctx) {
			this.Redirect("/login", 302)
			return
		}

		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return

	case "del":
		if !checkAccount(this.Ctx) {
			this.Redirect("/login", 302)
			return
		}

		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DeleteTopics(id)
		if err != nil {
			beego.Error(err)
		}

		err = models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	}

	this.TplNames = "category.html"
	this.Data["IsCategory"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	this.locale()
}
