package models

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "time"
)

type Article struct {
    Id            int         `form:"-"`
    Title         string      `orm:"size(255)" form:"title" valid:"Required;Range(2, 10)"`
    Description   string      `orm:"size(255)" form:"description" valid:"Required;Range(2, 10)"`
    Created       time.Time   `orm:"auto_now_add;type(datetime)" form:"-"`
    Updated       time.Time   `orm:"auto_now;type(datetime)" form:"-"`
}

func init(){
    orm.RegisterModel(new(Article))
}

func GetAllArticles() []*Article {
  o := orm.NewOrm()
  var articles []*Article
  o.QueryTable(new(Article)).All(&articles)
  return articles
}

func DeleteArticle(id int) bool {
  o := orm.NewOrm()
  _, err := o.Delete(&Article{Id: id})
  if err == nil {
    // successfully delete
    err := o.Read(&Article{Id: int(id)})
		if err == orm.ErrNoRows {
      // Obj not found
			return true
		}
    beego.Error(err)
    return false
  }

  beego.Error(err)
  return false
}

func GetArticle(id int) *Article {
  o := orm.NewOrm()
	article := Article{Id: id}
	o.Read(&article)
  return &article
}

func InsertArticle(article Article) bool {
	o := orm.NewOrm()
	id, err := o.Insert(&article)

  if err == nil {
		// successfully inserted
		if err := o.Read(&Article{Id: int(id)}); err == orm.ErrNoRows {
      // Obj not found
      beego.Error(err)
			return false
		}
    return true
	}

  beego.Error(err)
  return false
}

func UpdateArticle(article Article) bool {
	o := orm.NewOrm()
	a := Article{Id: article.Id}

	// get existing user
  err := o.Read(&a)
	if err != orm.ErrNoRows {
    // Obj found
		a.Title = article.Title
		a.Description = article.Description
		_, err := o.Update(&a)

		if err == nil {
			// update successful
      return true
		}
	}

  beego.Error(err)
	return false
}
