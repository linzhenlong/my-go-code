package main

import (
	"github.com/kataras/iris"
)

func main() {
	// 初始化框架
	app := iris.New()
	// 日志等级
	app.Logger().SetLevel("debug")
	// 注册模板目录
	app.RegisterView(iris.HTML("./web/views", "html"))

	// 注册控制器

	// 启动
	app.Run(
		iris.Addr(":8080"),
	)

}
