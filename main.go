package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"github.com/xuzhenglun/Blog-Go/controllers"
	"github.com/xuzhenglun/Blog-Go/models"
	"os"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.AutoRouter(&controllers.ReplyController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	os.Mkdir("attachment", os.ModePerm)
	beego.SetStaticPath("/attachment", "attachment")

	err := i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	if err != nil {
		beego.Error(err)
	}
	err = i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	if err != nil {
		beego.Error(err)
	}
	err = i18n.SetMessage("zh-HK", "conf/locale_zh-HK.ini")
	if err != nil {
		beego.Error(err)
	}

	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
