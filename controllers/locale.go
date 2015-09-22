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
	case "zh-HK":
		this.Data["Lang"] = "zh-HK"
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
	this.Data["category_lang"] = "Category"
	this.Data["created_lang"] = "Created"
	this.Data["replys_lang"] = "Replies"
	this.Data["views_lang"] = "Views"
	this.Data["topic_lang"] = "Topic"
	this.Data["home_lang"] = "Home"
	this.Data["tags_lang"] = "Tags"
	this.Data["contment_lang"] = "Content"
	this.Data["attach_lang"] = "Attachment"
	this.Data["login_lang"] = "Login"
	this.Data["logout_lang"] = "Logout"
	this.Data["cate_list_lang"] = "Category List"
	this.Data["cate_add_lang"] = "Add Category"
	this.Data["add_lang"] = "Add"
	this.Data["delete_lang"] = "Delete"
	this.Data["action_lang"] = "Action"
	this.Data["cate_count_lang"] = "Topic Count"
	this.Data["invild_cate_lang"] = "Invild Category"
	this.Data["topic_list_lang"] = "Topic List"
	this.Data["updated_lang"] = "Updated"
	this.Data["views_lang"] = "Views"
	this.Data["replies_lang"] = "Replies"
	this.Data["last_rep_lang"] = "Last Reply"
	this.Data["modify_lang"] = "Modify"
	this.Data["add_topic_lang"] = "Add Topic"
	this.Data["modity_topic_lang"] = "Modify Topic"
	this.Data["submit_lang"] = "Submit"
	this.Data["nickname_lang"] = "Nickname"
	this.Data["account_lang"] = "Account"
	this.Data["pwd_lang"] = "Password"
	this.Data["remember_lang"] = "Remember me"
	this.Data["signin_lang"] = "Sign in"
	this.Data["back_lang"] = "Back"
	this.Data["invild_act_lang"] = "Invild Username"
	this.Data["invild_pwd_lang"] = "Invild Password"
}

func (this *baseController) SetLang(lang string) {
	maxAge := 1<<31 - 1
	this.Ctx.SetCookie("Lang", lang, maxAge, "/")
	this.Redirect("/", 302)
}
