package controllers

import (
  "github.com/astaxie/beego"
  "test/models"
  "strconv"
  "encoding/json"
)

type ArticleController struct {
  beego.Controller
}

func (c *ArticleController) All(){
  c.Data["Articles"] = models.GetAll()
  c.TplName = "articles/index.tpl"
}
func (c *ArticleController) Add(){
  // var article models.Article
}
func (c *ArticleController) Delete(){
  id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
  models.DeleteArticle(id)
  c.TplName = "articles/delete.tpl"
}

func (c *ArticleController) Show(){
  id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
  c.Data["Article"] = models.GetArticle(id)
  c.TplName = "articles/show.tpl"
}

func (c *ArticleController) Insert() {
	var a models.Article
	json.Unmarshal(c.Ctx.Input.RequestBody, &a)
  c.Data["Article"] = models.InsertArticle(a)
  c.TplName = "articles/show.tpl"
}
