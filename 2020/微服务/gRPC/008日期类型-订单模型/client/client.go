package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/008日期类型-订单模型/client/services"
	"google.golang.org/grpc"
	"time"
)

func main() {
	clientConn, err := grpc.Dial(":6061", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer clientConn.Close()

	// 创建商品服务客户端
	serviceClient := services.NewProductServiceClient(clientConn)

	// 创建订单服务客户端
	orderServiceClient := services.NewOrderServiceClient(clientConn)

	// 获取商品列表
	ctx := context.Background()
	productRequest := &services.ProductRequest{
		ProId: 1,
		Size:  18,
	}
	productList, err := serviceClient.GetProductList(ctx, productRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(productList)

	orderMain := &services.OrderMain{
		OrderId:    1,
		OrderNo:    "101001",
		UserId:     1,
		OrderMoney: 180.00,
		OrderTime: &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
		},
	}
	orderResponse, err := orderServiceClient.NewOrder(ctx, orderMain)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(orderResponse)
}
