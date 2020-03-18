package main

import (
	"github.com/kataras/iris"
)

func main() {
	// 创建iris 实例
	app := iris.New()
	// 设置错误等级
	app.Logger().SetLevel("debug")

	// 注册模板
	tmplate := iris.HTML("./web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)

	// 设置静态文件目录
	app.StaticWeb("/assets", "./web/assets")

	// 异常跳转到指定错误页
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// 注册控制器，实现路由
	app.Get("/",func(ctx iris.Context) {
		ctx.HTML("hello")
	})

	// 启动服务
	app.Run(
		iris.Addr(":8081"),
		iris.WithConfiguration(iris.TOML("./configs/main.tml")),
	)

}
