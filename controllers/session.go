package controllers

import (
	// "fmt"
	"home/models"

	// "github.com/astaxie/beego/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

//SessionController jiegouti jicheng beego.controller
type SessionController struct {
	beego.Controller
}

//RetData fangfa
func (c *SessionController) RetData(resp interface{}) {
	//给客户端返回json数据
	c.Data["json"] = resp //将数据写入data[json]中
	//将json写回客户端
	c.ServeJSON() //调用此方法进行渲染，json数据即可序列化输出
}

//GetSessionData 重写get方法
func (c *SessionController) GetSessionData() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	user := models.User{}
	// resp["errno"] = 0
	// resp["errmsg"] = "err"
	name := c.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = 0
		resp["errmsg"] = "ok"
		resp["data"] = user
	}
	//c.RetData(resp)
}
//DeleteSessionDate fangfa
func (c *SessionController) DeleteSessionDate() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	err := c.DelSession("name")
	if err!= nil{
		resp["errno"] = 1
		resp["errmsg"] = "失败"
	}
	resp["errno"] = 0
	resp["errmsg"] = "成功"
	
}