package controllers

import (
	// "fmt"
	// "home/models"

	// "github.com/astaxie/beego/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

//HouseIndexController jiegouti jicheng beego.controller
type HouseIndexController struct {
	beego.Controller
}
//RetData fangfa
func (c *HouseIndexController) RetData(resp interface{}) {
	//给客户端返回json数据
	c.Data["json"] = resp //将数据写入data[json]中
	//将json写回客户端
	c.ServeJSON() //调用此方法进行渲染，json数据即可序列化输出
}

//GetHouseIndex 重写get方法
func (c *HouseIndexController) GetHouseIndex() {
	resp := make(map[string]interface{})
	resp["errno"] = 0
	resp["errmsg"] = "查询成功"
	c.RetData(resp)
}