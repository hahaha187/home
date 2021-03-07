package controllers

import (
	"fmt"
	"home/models"

	"github.com/astaxie/beego/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

//AreaController jiegouti jicheng beego.controller
type AreaController struct {
	beego.Controller
}

//RetData fangfa
func (c *AreaController) RetData(resp interface{}) {
	//给客户端返回json数据
	c.Data["json"] = resp //将数据写入data[json]中
	//将json写回客户端
	c.ServeJSON() //调用此方法进行渲染，json数据即可序列化输出
}

//GetArea 重写get方法
func (c *AreaController) GetArea() {
	fmt.Println("connect success")
	resp := make(map[string]interface{}) //map需要用make方法分配确定的内存地址，否则网页返回assignment to entry in nil map
	//defer c.RetData(resp)
	//从session拿数据
	//1.从数据库拿到数据
	var areas []models.Area
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	fmt.Println(num)
	if err != nil {
		resp["errno"] = 4001
		resp["errmsg"] = "数据库查询错误"
		c.RetData(resp)
		return
	}
	if num == 0 {
		resp["errno"] = 4002
		resp["errmsg"] = "无数据"
		c.RetData(resp)
		return
	}
	resp["errno"] = 0
	resp["errmsg"] = "成功"
	resp["data"] = areas
	//2.打包成json返回给前端
	c.RetData(resp)
}