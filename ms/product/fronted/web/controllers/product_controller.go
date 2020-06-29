package controllers

import (
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"github.com/linzhenlong/my-go-code/ms/product/services"
)

// ProductController .
type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
	OrderService   services.IOrderService
	Sesssion       *sessions.Session
}

var (
	// 生成html保存目录
	htmlOutPath = "./web/htmlProductShow/"
	// 静态文件模板目录
	templatePath = "./web/views/template/"
)

//generateStaticHtml 生成html静态文件.
func generateStaticHtML(ctx iris.Context, template *template.Template, fileName string, product *datamodels.Product) {
	// 1.判断静态文件生成存在.
	if existFile(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			ctx.Application().Logger().Debug(err)
		}
	}
	// 2.生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		ctx.Application().Logger().Debug(err)
	}
	defer file.Close()
	ctx.Application().Logger().Debug("generateStaticHtML--->", product)
	template.Execute(file, &product)
}

// existFile 判断文件是否存在.
func existFile(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func (p *ProductController) GetGenHtml() {
	contentTemp, err := template.ParseFiles(filepath.Join(templatePath, "product.html"))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")

	id, err := p.Ctx.URLParamInt64("id")
	if err != nil || id < 0 {
		//todo xxx
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		//todo xxx
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Application().Logger().Debug(product)

	// 生成静态文件
	generateStaticHtML(p.Ctx, contentTemp, fileName, product)
	p.Ctx.Application().Logger().Debug(contentTemp)

}

// GetDetail 文章详情页.
func (p *ProductController) GetDetail() mvc.View {
	id, err := p.Ctx.URLParamInt64("id")
	if err != nil || id < 0 {
		//todo xxx
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		//todo xxx
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Application().Logger().Debug(product)

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

// GetOrder .
func (p *ProductController) GetOrder() mvc.View {
	productID, _ := p.Ctx.URLParamInt("productID")
	userIDstring := p.Ctx.GetCookie("uid")

	userID, _ := strconv.Atoi(userIDstring)

	var orderID int64
	var message = "抢购失败"
	// 获取商品详情
	productInfo, err := p.ProductService.GetProductByID(int64(productID))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
		goto RES
	}
	// 判断商品数量
	if productInfo.ProductNum > 0 {

		// 创建订单
		order := &datamodels.Order{
			UserID:      int64(userID),
			ProductID:   int64(productID),
			OrderStatus: datamodels.OrderSuccess,
		}
		orderID, err = p.OrderService.InsertOrder(order)
		if err != nil {
			p.Ctx.Application().Logger().Debug("InsertOrder error:", err)
			message = err.Error()
			goto RES
		}

		// 扣除商品数量 -1
		productInfo.ProductNum--
		err = p.ProductService.UpdateProduct(productInfo)
		if err != nil {
			p.Ctx.Application().Logger().Debug("UpdateProduct error:", err)
			message = err.Error()
			goto RES
		}
		message = "抢购成功啦....."
	}
RES:
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			"orderID": orderID,
			"message": message,
		},
	}
}
