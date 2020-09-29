package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/linzhenlong/my-go-code/my-go-frame/app"
)

func main() {
	flag.Parse()

	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Gin.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(
			200,
			gin.H{
				"message": "success",
			},
		)
	})
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
