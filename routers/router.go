package routers

import (
	"test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/articles", &controllers.ArticleController{}, "get:List")
    beego.Router("/articles/:id([0-9]+)", &controllers.ArticleController{}, "get:Show")
    beego.Router("/articles/:id([0-9]+)/delete", &controllers.ArticleController{}, "get,delete:Delete")
    beego.Router("/articles/:id([0-9]+)/edit", &controllers.ArticleController{}, "get,post:Update")
    beego.Router("/articles/save", &controllers.ArticleController{}, "get,post:Save")
}
