package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"tiktop/db"
	"tiktop/midware"
	"tiktop/views"
)

func main() {
	// 初始化数据库
	err := db.Init_db()
	//出错就结束
	if err != nil {
		os.Exit(-1)
	}
	//启动 gin路由
	r := gin.Default()
	auth := midware.Jwt2r()
	//用啦存放静态文件的
	r.StaticFS("/static", http.Dir("./static"))
	//设置组
	Douyin_router := r.Group("/douyin")

	//下面分别是基本API
	//1.登录相关
	Douyin_router.GET("/user/", auth, views.UserInfo)
	Douyin_router.POST("/user/register/", views.Register)
	Douyin_router.POST("/user/login/", views.Login)
	//2.视频流相关

	//互动API

	//社交接口（尽量完成）

	r.Run()
}
