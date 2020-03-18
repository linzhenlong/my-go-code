package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/linzhenlong/my-go-code/ms/Iris/web/controllers"
)

func main() {
	// 初始化框架
	app := iris.New()
	// 日志等级
	app.Logger().SetLevel("info")
	// 注册模板目录
	app.RegisterView(iris.HTML("./web/views/", ".html"))

	// 注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	// 启动
	app.Run(
		iris.Addr(":8081"),
	)

}
