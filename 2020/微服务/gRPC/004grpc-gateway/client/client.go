package main

import (
	"context"
	"fmt"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/004grpc-gateway/client/helper"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/004grpc-gateway/client/services"
	"google.golang.org/grpc"
	"log"
)

func main() {

	clientConn, err := grpc.Dial(":6061",grpc.WithTransportCredentials(helper.GetClientCred()))
	if err != nil {
		log.Fatal(err)
	}
	defer clientConn.Close()

	// 初始化客户端
	productServiceClient := services.NewProductServiceClient(clientConn)

	searchServiceClient := services.NewSearchServiceClient(clientConn)



	// 获取库存
	ctx := context.Background()
	productRequest := &services.ProductRequest{}
	stockStatus, err := productServiceClient.GetProductStockStatus(ctx, productRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("库存是",stockStatus.StockStatus)

	productName, err := productServiceClient.GetProductName(ctx, productRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("商品名称是",productName.ProductName)

	// 文章列表
	articles, err := searchServiceClient.GetArticles(ctx, &services.SearchRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("articles",articles)

}
