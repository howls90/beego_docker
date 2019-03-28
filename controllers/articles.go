package controllers

import (
  "github.com/astaxie/beego"
  "test/models"
  "strconv"
  "github.com/astaxie/beego/validation"
)

type ArticlesController struct {
  ExtendedController
}

// articles/
func (this *ArticlesController) List(){
  beego.ReadFromRequest(&this.Controller)
  this.Data["Articles"] = models.GetAllArticles()
  this.TplName = "articles/index.tpl"
}

// articles/save
func (this *ArticlesController) Save(){
  if this.Ctx.Input.Method() == "POST" {
    at := ArticleValidate(this)
    res := models.InsertArticle(at)
    ArticleFlash(this, "Article saved!", res)
    this.Ctx.Redirect(302, "/articles")
  } else {
    this.TplName = "articles/save.tpl"
  }
}
// articles/:id/delete
func (this *ArticlesController) Delete(){
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  res := models.DeleteArticle(id)
  ArticleFlash(this, "Article deleted!", res)

  this.Ctx.Redirect(302, "/articles")
}

// articles/:id/show
func (this *ArticlesController) Show(){
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  this.Data["Article"] = models.GetArticle(id)
  this.TplName = "articles/show.tpl"
}

// articles/:id/edit
func (this *ArticlesController) Update(){
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  if this.Ctx.Input.Method() == "POST" {
    at := ArticleValidate(this)
    at.Id = id
    res := models.UpdateArticle(at)
    ArticleFlash(this, "Article edited!", res)
  } else {
    this.Data["Article"] = models.GetArticle(id)
    this.TplName = "articles/edit.tpl"
  }
}

/*
Parser and validate the request post data

@param controller

@return article
*/
func ArticleValidate(this *ArticlesController) models.Article{
  at := models.Article{}
  if err := this.ParseForm(&at); err != nil {
    beego.Error(err)
    this.Abort("500")
  }
  // errors
  valid := validation.Validation{}
  res, err := valid.Valid(&at)
  if err != nil {
		beego.Error(err)
	}
	if !res {
		for _, err := range valid.Errors {
        beego.Info(err.Key, err.Message)
    }
	}
  
  return at
}

/*
Check if the last statement doesn't have erros and return flash message

@param controller
@param flash message
@param last statement (CRUD)

*/
func ArticleFlash(this *ArticlesController, msn string, res bool){
  if (res == true){
    flash := beego.NewFlash()
    flash.Notice(msn)
    flash.Store(&this.Controller)
    this.Ctx.Redirect(302, "/articles")
  } else {
    beego.Error(res)
    this.Abort("500")
  }
}
