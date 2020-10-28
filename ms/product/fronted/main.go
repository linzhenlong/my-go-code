package main

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/fronted/middleware"
	"github.com/linzhenlong/my-go-code/ms/product/fronted/web/controllers"
	"github.com/linzhenlong/my-go-code/ms/product/repositories"
	"github.com/linzhenlong/my-go-code/ms/product/services"
	"log"
	"time"
)

func main() {
	// 创建iris实例
	app := iris.New()

	// 日志等级.
	app.Logger().SetLevel("debug")

	//注册模板
	tmplate := iris.HTML("./web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)

	// 设置静态文件目录
	app.StaticWeb("/public", "./web/public")
	app.StaticWeb("/html", "./web/htmlProductShow")

	// 异常指定错误页
	app.OnAnyErrorCode(func(ctx iris.Context) {
		log.Printf("%v", ctx.Values().GetStringDefault("message", "访问页面出错"))
		log.Printf("%v", ctx.GetStatusCode())
		//ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问页面出错"))
		ctx.ViewData("message", ctx.GetStatusCode())
		ctx.ViewData("message2", "访问页面出错")
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	// 数据库连接.
	gormDB, err := common.NewGorm()
	if err != nil {
		log.Fatalf("mysql-gorm连接error:%v", err)
	}
	mysqlDB, err := common.NewMysqlConn()
	if err != nil {
		log.Fatalf("NewMysqlConn连接error:%v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// session

	sess := sessions.New(sessions.Config{
		Cookie:  "hello-world",
		Expires: 60 * time.Minute,
	})

	// 注册控制器
	userRpository := repositories.NewUserRepository("user_ms", gormDB)
	userService := services.NewUserService(userRpository)

	userParty := app.Party("/user")
	user := mvc.New(userParty)
	user.Register(ctx, userService, sess.Start)
	user.Handle(new(controllers.UserController))

	// 注册商品控制器
	productRepository := repositories.NewProductManager("product", mysqlDB, gormDB)
	productService := services.NewProductService(productRepository)

	// 注册order
	orderRepository := repositories.NewOrderManager("order", gormDB)
	orderService := services.NewOrderService(orderRepository)

	productParty := app.Party("/product")
	product := mvc.New(productParty)

	// 中间件
	productParty.Use(middleware.AuthProductController)
	product.Register(ctx, productService, sess.Start, orderService)
	product.Handle(new(controllers.ProductController))

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("hello iris")
	})
	// 启动服务
	app.Run(
		iris.Addr(":8082"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
