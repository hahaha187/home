package controllers

import (
	"fmt"
	"home/models"

	"github.com/astaxie/beego/client/orm"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

//UserController jiegouti jicheng beego.controller
type UserController struct {
	beego.Controller
}

//RetData fangfa
func (c *UserController) RetData(resp interface{}) {
	//给客户端返回json数据
	c.Data["json"] = resp //将数据写入data[json]中
	//将json写回客户端
	c.ServeJSON() //调用此方法进行渲染，json数据即可序列化输出
}
//Reg fangfa 
func (c *UserController) Reg() {
	resp := make(map[string] interface{})
	defer c.RetData(resp)
	//获取前端传来的json数据
	json.Unmarshal(c.Ctx.Input.RequestBody,&resp)
	// fmt.Println(resp["mobile"])  // 验证能否拿到前端传来的数据
	// fmt.Println(resp["password"])
	// fmt.Println(resp["sms_code"])

	//插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Password = resp["password"].(string) //resp类型是interface,而password的类型是string，需要断言 
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)
	id,err := o.Insert(&user)
	if err != nil {
		resp["errno"] = 4002
		resp["errmsg"] = "注册失败"
		return
	}
	//fmt.Printf("注册成功，id :",id)
	fmt.Println(id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"
	c.SetSession("name",user.Name)
	// resp["errno"] = 4001
	// resp ["errmsg"] = "查询失败"
	// c.RetData(resp)
}