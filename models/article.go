package models

import (
  "github.com/astaxie/beego/orm"
  "time"
  "fmt"
)

type Article struct {
    Id            int         
    Title         string      `orm:null`
    Description   string      `orm:null`
    Date          time.Time   `omr:"auto_now_add;type(datetime)"`

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
    // successfull
    return true
  }

  return false
}

func GetArticle(id int) *Article {
  o := orm.NewOrm()
	article := Article{Id: id}
	o.Read(&article)
  return &article
}

func InsertArticle(article Article) *Article {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))

	// get prepared statement
	i, _ := qs.PrepareInsert()

	var a Article
	article.Date = time.Now()

	// Insert
	id, err := i.Insert(&article)
  if err == nil {
		// successfully inserted
		a = Article{Id: int(id)}
		err := o.Read(&a)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

  return &a
}

func UpdateArticle(article Article) *Article {
	o := orm.NewOrm()
	oldArticle := Article{Id: article.Id}
  fmt.Println("dins")
	// var newArticle Article

	// get existing user
	// if o.Read(&oldArticle) == nil {

    o.Read(&oldArticle)
		oldArticle.Title = article.Title
		oldArticle.Description = article.Description
		oldArticle.Date = time.Now()
		_, err := o.Update(&oldArticle)

		// read updated user
		if err == nil {
			// update successful
			// newArticle = Article{Id: article.Id}
			// o.Read(&newArticle)
		}
	// }

	return &oldArticle
}
