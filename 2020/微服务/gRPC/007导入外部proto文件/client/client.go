package main

import (
	"context"
	"fmt"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/007导入外部proto文件/server/services"
	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial(":6061", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer clientConn.Close()

	// 创建客户端
	serviceClient := services.NewProductServiceClient(clientConn)

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
}
