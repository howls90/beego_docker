package models

import (
  "github.com/astaxie/beego/orm"
  "time"
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
