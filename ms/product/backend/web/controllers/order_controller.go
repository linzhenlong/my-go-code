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

// OrderController 订单控制器.
type OrderController struct {
	Ctx            iris.Context
	OrderService   services.IOrderService
	ProductService services.IProductService
}

const pageSize = 5

// ProductOrder 商品订单
type ProductOrder struct {
	ProductID    int64
	ProductName  string
	ProductNum   int64
	ProductImage string
	ProductURL   string
	OrderID      int64
	UserID       int64
	OrderStatus  int64
}

// GetEdit 编辑页
func (o *OrderController) GetEdit() mvc.View {
	id, err := o.Ctx.URLParamInt64("id")
	orderInfo, err := o.OrderService.GetOrderByID(id)
	if id <= 0 || err != nil {
		o.Ctx.Redirect("/order/list")
	}
	orderProduct := ProductOrder{}
	product, err := o.ProductService.GetProductByID(orderInfo.ProductID)
	orderProduct.ProductID = product.ID
	orderProduct.ProductName = product.ProductName
	orderProduct.ProductNum = product.ProductNum
	orderProduct.ProductImage = product.ProductImage
	orderProduct.ProductURL = product.ProductURL
	orderProduct.OrderID = orderInfo.ID
	orderProduct.UserID = orderInfo.UserID
	orderProduct.OrderStatus = orderInfo.OrderStatus

	log.Printf("%#v", orderProduct)

	return mvc.View{
		Name: "order/edit",
		Data: iris.Map{
			"order": orderProduct,
		},
	}
}

// PostUpdate 订单更新.
func (o *OrderController) PostUpdate() {
	order := datamodels.Order{}
	err := o.Ctx.ReadForm(&order)
	if err != nil && !iris.IsErrPath(err) {
		o.Ctx.StatusCode(iris.StatusInternalServerError)
		o.Ctx.HTML(err.Error())
		log.Printf("%#v", err)
	}
	log.Printf("%#v", order)
	err = o.OrderService.UpdateOrder(&order)
	if err != nil {
		o.Ctx.HTML("<script>alert('err');location.href='/order/list';</script>")
	} else {
		o.Ctx.HTML("<script>alert('success');location.href='/order/list';</script>")
	}
}

// GetList 订单列表 .
func (o *OrderController) GetList() mvc.View {
	params := make(map[string]interface{})
	page, err := o.Ctx.URLParamInt("page")
	if err != nil || page < 0 {
		page = 1
	}

	urlPre := "/order/list?"
	offset := (page - 1) * pageSize

	orderStatus, _ := o.Ctx.URLParamInt64("order_status")

	params["offset"] = offset
	if orderStatus >= 0 {
		urlPre += "order_status=" + strconv.Itoa(int(orderStatus))
		params["orderStatus"] = orderStatus
	}
	params["limit"] = pageSize
	orders, err := o.OrderService.GetOrderList(params)
	if err != nil {
		log.Printf("err:%v", err)
	}
	list := []ProductOrder{}
	for _, v := range orders {
		orderProduct := ProductOrder{}
		orderProduct.OrderID = v.ID
		orderProduct.UserID = v.UserID
		orderProduct.OrderStatus = v.OrderStatus
		orderProduct.ProductID = v.ProductID
		product, err := o.ProductService.GetProductByID(orderProduct.ProductID)
		if err != nil {
			continue
		}
		orderProduct.ProductName = product.ProductName
		orderProduct.ProductNum = product.ProductNum
		orderProduct.ProductImage = product.ProductImage
		list = append(list, orderProduct)
	}
	total := o.OrderService.GetOrderTotal(params)

	pageData := common.Paginator(page, pageSize, total)
	for k, v := range pageData {
		log.Printf("k=>%v,v:%v", k, v)
	}

	return mvc.View{
		Name: "order/list",
		Data: iris.Map{
			"orderList":  list,
			"orderTotal": total,
			"pageData":   pageData,
			"Url":        urlPre,
		},
	}
}

// GetDel 删除订单.
func (o *OrderController) GetDel() {
	id, err := o.Ctx.URLParamInt64("id")
	if err != nil || id < 0 {
		o.Ctx.HTML("<script>alert('err');location.href='/order/list';</script>")
	}
	status := o.OrderService.DeleteOrder(id)
	if status {
		o.Ctx.HTML("<script>alert('success');location.href='/order/list';</script>")
	} else {
		o.Ctx.HTML("<script>alert('err');location.href='/order/list';</script>")
	}
}
