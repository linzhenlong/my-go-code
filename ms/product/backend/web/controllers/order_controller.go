package controllers

import (
	"log"
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/linzhenlong/my-go-code/ms/product/common"
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
	return mvc.View{
		Name: "order/edit",
		Data: iris.Map{
			"order": orderInfo,
		},
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
