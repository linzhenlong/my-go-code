package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"hello golang",
	})
}

func main() {
	// 返回默认路由引擎.
	r := gin.Default()

	// 指定用户使用get方式请求
	r.GET("/hello",sayHello)

	// 获取一本书
	r.GET("/book",func(c *gin.Context){
		c.JSON(200, gin.H{
			"msg":"GET",
		})
	})
	r.POST("/book",func(c *gin.Context){
		c.JSON(200, gin.H{
			"msg" : "POST",
		})
	})

	r.PUT("/book",func(c *gin.Context){
		c.JSON(200,gin.H{
			"msg":"PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"msg":"DELETE",
		})
	})

	r.Any("/any",func(c *gin.Context){
		c.JSON(200, gin.H{
			"msg":c.Request.Method,
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	err := r.Run(":8888")
	if err != nil {
		fmt.Printf("gin run error=%v", err)
	}
}