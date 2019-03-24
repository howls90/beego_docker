package controllers

import (
  "github.com/astaxie/beego"
)

type ExtendedController struct {
  beego.Controller
}

func (this *ExtendedController) Prepare() {
  this.Layout = "layout.html"
}
