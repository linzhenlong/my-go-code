package main

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/linzhenlong/my-go-code/ms/product/backend/web/controllers"
	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/repositories"
	"github.com/linzhenlong/my-go-code/ms/product/services"
	"log"
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
		log.Printf("%v", ctx.Values().GetStringDefault("message", "访问页面出错"))
		log.Printf("%v", ctx.GetStatusCode())
		//ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问页面出错"))
		ctx.ViewData("message", ctx.GetStatusCode())
		ctx.ViewData("message2", "访问页面出错")
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	// 链接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Fatalf("mysql 连接error:%v", err)
	}

	// 创建上下文环境
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 注册控制器，实现路由
	productRpository := repositories.NewProductManager("product", db)
	prodcutService := services.NewProductService(productRpository)

	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, prodcutService)
	product.Handle(new(controllers.ProductController))

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("hello")
	})

	// 启动服务
	app.Run(
		iris.Addr(":8081"),
		iris.WithConfiguration(iris.TOML("./configs/main.tml")),
	)

}
