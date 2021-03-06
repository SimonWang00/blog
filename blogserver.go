package main

//File  : blogserver.go
//Author: Simon
//Describe: application
//Date  : 2020/12/8

import (
	"gin.blog/config"
	"gin.blog/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

func main() {
	go startWebServer()		// 启动web服务
	// 阻塞等子线程，Block...
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

// hello 调试请求
func hello(c *gin.Context)  {
	c.String(http.StatusOK,"Welcom to my gin.blog!")
}

// startWebServer 启动web服务
func startWebServer() {
	// setup my apps routers
	g := routes.SetRegisterRouters()
	// 静态资源加载，本项目css,js以及资源图片
	g.StaticFS("/assets/blog", http.Dir("./public/assets/blog"))
	g.StaticFS("/assets/admin", http.Dir("./public/assets/admin"))
	g.LoadHTMLGlob("public/views/**/*")
	// 启动 http server
	if err := http.ListenAndServe(config.AppConfig.Addr, g); err != nil {
		log.Fatal("http server 启动失败", err)
	}
}