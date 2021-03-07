package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)
//MainController jiegouti
type MainController struct {
	beego.Controller
}
//Get fangfachongxie
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
