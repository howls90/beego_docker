package routers

import (
	"test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/articles", &controllers.ArticleController{}, "get:All")
    beego.Router("/articles/:id", &controllers.ArticleController{}, "get:Show")
    beego.Router("/articles/add", &controllers.ArticleController{}, "put:Insert")
    beego.Router("/articles/delete/:id", &controllers.ArticleController{}, "get:Delete")
}
