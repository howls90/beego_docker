package routers

import (
	"test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/articles", &controllers.ArticlesController{}, "get:List")
    beego.Router("/articles/:id([0-9]+)", &controllers.ArticlesController{}, "get:Show")
    beego.Router("/articles/:id([0-9]+)/delete", &controllers.ArticlesController{}, "get,delete:Delete")
    beego.Router("/articles/:id([0-9]+)/edit", &controllers.ArticlesController{}, "get,post:Update")
    beego.Router("/articles/save", &controllers.ArticlesController{}, "get,post:Save")
}
