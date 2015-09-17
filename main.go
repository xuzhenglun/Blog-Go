package main

import (
	"blog/controllers"
	"blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	beego.Run()
}
