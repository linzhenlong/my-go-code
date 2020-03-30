package controllers

import (
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"github.com/linzhenlong/my-go-code/ms/product/services"
)

// ProductController 商品控制器.
type ProductController struct {
	Ctx            iris.Context // 上下文环境.
	ProductService services.IProductService
}

//GetList 商品列表
func (p *ProductController) GetList() mvc.View {
	productList, err := p.ProductService.GetAllProduct()

	/* for i, v := range productList {
		log.Printf("%d =>%v", i, v.ID)
		log.Printf("%d =>%v", i, v.ProductName)
	} */

	if err != nil {
		log.Printf("err:%v", err)
	}
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productList": productList,
		},
	}
}

func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}

	err := p.Ctx.ReadForm(product)
	if err != nil {
		p.Ctx.StatusCode(iris.StatusInternalServerError)
		p.Ctx.WriteString(err.Error())
	}
	p.Ctx.Application().Logger().Debugf("%#v", product)
	//p.Ctx.Writef("%#v", product)
	err = p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.StatusCode(iris.StatusInternalServerError)
		p.Ctx.WriteString(err.Error())
	}
	p.Ctx.Redirect("/product/list")
}
