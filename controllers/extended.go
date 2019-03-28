package controllers

import (
  "github.com/astaxie/beego"
  "html/template"
)

type ExtendedController struct {
  beego.Controller
}

func (this *ExtendedController) Prepare() {
  this.Layout = "layout.html"
  this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
}
