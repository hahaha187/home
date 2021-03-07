package main

import (
	_ "home/models"
	_ "home/routers"
	"log"
	"net/http"
	"strings"

	//"github.com/astaxie/beego/server/web"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func main() {
	ignoreStaticPath()
	beego.Run()
}

//重定向static静态路径
func ignoreStaticPath() {

	//透明static

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

// TransparentStatic ff
func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	log.Println(orpath)
	//如果url中有api，则取消静态路由重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+orpath)
	//将全部的静态资源重定向 加上/static/html路径
	//http://ip:port:8080/index.html----> http://ip:port:8080/static/html/index.html
	//如果restFUL api  那么就取消冲定向
	//http://ip:port:8080/api/v1.0/areas ---> http://ip:port:8080/static/html/api/v1.0/areas
}
