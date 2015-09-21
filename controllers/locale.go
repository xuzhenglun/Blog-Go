package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare() {
	lang := this.GetString("lang")

	switch lang {
	case "zh-CN":
		this.Data["Lang"] = "zh-CN"
	case "en-US":
	default:
		this.Data["Lang"] = "en-US"
		return
	}

	if len(lang) > 0 {
		this.SetLang(lang)
	}
}

func (this *baseController) locale() {
	lang_Cookie, err := this.Ctx.Request.Cookie("Lang")
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Lang"] = lang_Cookie.Value
	}

	this.Data["title_lang"] = "Title"
	this.Data["cate_lang"] = "Categories"
	this.Data["created_lang"] = "Created"
	this.Data["replys_lang"] = "Replys"
	this.Data["views_lang"] = "Views"
	this.Data["topic_lang"] = "Topic"
	this.Data["home_lang"] = "Home"
	this.Data["login_lang"] = "Login"
	this.Data["logout_lang"] = "Logout"
}

func (this *baseController) SetLang(lang string) {
	maxAge := 1<<31 - 1
	this.Ctx.SetCookie("Lang", lang, maxAge, "/")
	this.Redirect("/", 302)
}
