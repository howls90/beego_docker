package controllers

import (
  "github.com/astaxie/beego"
  "test/models"
  "strconv"
  // "fmt"
)

type ArticleController struct {
  ExtendedController
}

func (this *ArticleController) List(){
  beego.ReadFromRequest(&this.Controller)
  this.Data["Articles"] = models.GetAllArticles()
  this.TplName = "articles/index.tpl"
}

func (this *ArticleController) Save(){
  if this.Ctx.Input.Method() == "POST" {
    var at models.Article
    at.Title = this.GetString("title")
    at.Description = this.GetString("description")
    res := models.InsertArticle(at)

    if (res == true){
      flash := beego.NewFlash()
      flash.Notice("Article saved!")
      flash.Store(&this.Controller)
    } else {
      beego.Error("Content bad format!")
      this.Abort("500")
    }
    this.Ctx.Redirect(302, "/articles")

  } else {
    this.TplName = "articles/save.tpl"
  }
}

func (this *ArticleController) Delete(){
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  res := models.DeleteArticle(id)
  if (res == true){
    flash := beego.NewFlash()
    flash.Notice("Article deleted!")
    flash.Store(&this.Controller)
  } else {
    beego.Error("Article not found!")
    this.Abort("500")
  }

  this.Ctx.Redirect(302, "/articles")
}

func (this *ArticleController) Show(){
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  this.Data["Article"] = models.GetArticle(id)
  this.TplName = "articles/show.tpl"
}

func (this *ArticleController) Update(){
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  if this.Ctx.Input.Method() == "POST" {
    var at models.Article
    at.Id = id
    at.Title = this.GetString("title")
    at.Description = this.GetString("description")
    models.UpdateArticle(at)

    flash := beego.NewFlash()
    flash.Notice("Article edit!")
    flash.Store(&this.Controller)
    this.Ctx.Redirect(302, "/articles")

  } else {
    this.Data["Article"] = models.GetArticle(id)
    this.TplName = "articles/edit.tpl"
  }
}
