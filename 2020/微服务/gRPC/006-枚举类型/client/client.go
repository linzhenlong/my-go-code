package main

import (
	"context"
	"fmt"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/006-枚举类型/server/services"
	"google.golang.org/grpc"
	"log"
)

func main() {
	clientConn, err := grpc.Dial(":6061", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer clientConn.Close()

	// 创建客户端
	serviceClient := services.NewProductServiceClient(clientConn)

	// 调用方法
	ctx := context.Background()
	productRequest := &services.ProductRequest{
		ProName: "xx",
		Size:    10,
		ProId:   1,
	}
	productList, err := serviceClient.GetProductList(ctx, productRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(productList)

	productInfo, err := serviceClient.GetProductInfo(ctx, &services.ProductRequest{
		ProArea: services.ProductAreas_BEI_JING,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(productInfo)

}
