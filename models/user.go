package models

import "github.com/astaxie/beego/orm"

type User struct {
    Id   int
	Name string
	Age  int
}

func init(){
    orm.RegisterModel(new(User))
}