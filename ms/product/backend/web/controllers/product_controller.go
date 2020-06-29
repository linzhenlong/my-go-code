package controllers

import (
	"log"
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/linzhenlong/my-go-code/ms/product/common"
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
	params := make(map[string]interface{})

	page, err := p.Ctx.URLParamInt("page")
	if err != nil || page < 0 {
		page = 1
	}

	urlPre := "/product/list?"
	offset := (page - 1) * pageSize
	params["offset"] = offset
	params["limit"] = pageSize
	productList, err := p.ProductService.SelectAllByParams(params)
	total := p.ProductService.GetTotal(params)

	/* for i, v := range productList {
		log.Printf("%d =>%v", i, v.ID)
		log.Printf("%d =>%v", i, v.ProductName)
	} */

	if err != nil {
		log.Printf("err:%v", err)
	}
	pageData := common.Paginator(page, pageSize, total)
	for k, v := range pageData {
		log.Printf("k=>%v,v:%v", k, v)
	}
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productList": productList,
			"total":       total,
			"pageData":    pageData,
			"Url":         urlPre,
		},
	}
}

// PostUpdate 更新商品
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

// PostAdd 添加商品.
func (p *ProductController) PostAdd() {
	product := &datamodels.Product{}
	err := p.Ctx.ReadForm(product)
	if err != nil {
		p.Ctx.StatusCode(iris.StatusInternalServerError)
		p.Ctx.WriteString(err.Error())
	}
	p.Ctx.Application().Logger().Debugf("%#v", product)
	_, err = p.ProductService.InsertProduct(product)
	if err != nil {
		p.Ctx.StatusCode(iris.StatusInternalServerError)
		p.Ctx.WriteString(err.Error())
	}
	p.Ctx.Redirect("/product/list")
}

//GetAdd 商品添加页
func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

// GetEdit 商品编辑.
func (p *ProductController) GetEdit() mvc.View {
	id := p.Ctx.URLParam("id")
	p.Ctx.Application().Logger().Debugf("urlid:%v", id)
	productID, _ := strconv.ParseInt(id, 10, 64)
	p.Ctx.Application().Logger().Debugf("%d", productID)
	product, err := p.ProductService.GetProductByID(productID)
	if err != nil {
		p.Ctx.StatusCode(iris.StatusInternalServerError)
		p.Ctx.WriteString(err.Error())
	}
	p.Ctx.Application().Logger().Debugf("prouduct:%#v", product)
	p.Ctx.Application().Logger().Debugf("err:%#v", err)
	return mvc.View{
		Name: "product/edit.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

// GetDel 删除一个商品
func (p *ProductController) GetDel() {
	id, _ := p.Ctx.URLParamInt64("id")
	p.Ctx.Application().Logger().Debugf("urlid:%v", id)
	status := p.ProductService.DeleteProductByID(id)
	if status {

		p.Ctx.WriteString("<script>alert('succ');location.href='/product/list';</script>")
	} else {
		p.Ctx.WriteString("<script>alert('err');location.href='/product/list';</script>")
	}

}
