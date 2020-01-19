package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func sayIndex(c *gin.Context) {
	
	c.HTML(http.StatusOK,"home.tmpl",gin.H{
		"title":"gin 渲染/templates/index.tmpl",
	})
}

func sayPosts(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl",gin.H{
		"title":"gin 渲染/templates/posts/index.tmpl",
	})
}
func sayUsers(c *gin.Context) {
	c.HTML(http.StatusOK,"users/1.tmpl", gin.H{
		"title":"gin 渲染 /templates/users/index.tmpl",
	})
}
func main() {
	router := gin.Default()
	
	// 加载单个模板文件
	//router.LoadHTMLFiles("./templates/posts/index.tmpl","./templates/users/index.tmpl")

	// 加载多个.
	router.LoadHTMLGlob("./templates/**/*")
	router.GET("/index", sayIndex)
	router.GET("/posts/index", sayPosts)
	router.GET("/users/index", sayUsers)
	router.Run(":9000")
}

//  https://www.bilibili.com/video/av79671612/ 18:00