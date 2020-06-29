package middleware

import (
	"github.com/kataras/iris"
)

// AuthProductController .
func AuthProductController(ctx iris.Context) {
	uid := ctx.GetCookie("uid")
	if uid == "" {
		ctx.Application().Logger().Debug("请先登录")
		uri := ctx.Request().RequestURI
		/* productID := ctx.URLParam("id")
		jumpURL := uri + "?id=" + productID */
		url := "/user/login?url=" + uri
		ctx.Application().Logger().Debug("要跳转回来", url)
		ctx.Redirect(url)
		return
	}
	ctx.Application().Logger().Debug("已登录uid:", uid)
	// 继续请求上下文
	ctx.Next()
}
