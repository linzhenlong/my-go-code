package main

import (
	"context"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/002/server/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {

	// 加载证书
	clientTLSFromFile, err := credentials.NewClientTLSFromFile("keys/server.crt", "linzl.com")
	if err != nil {
		log.Fatal(err)
	}
	// 创建grpc 链接
	clientConn, err := grpc.Dial(":6060",grpc.WithTransportCredentials(clientTLSFromFile))
	if err != nil {
		log.Fatal(err)
	}
	defer clientConn.Close()

	// 创建product service 客户端
	productServiceClient := services.NewProductServiceClient(clientConn)

	// 调用方法
	productRequest := &services.ProductRequest{
		ProId: 1,
	}
	productResponse, err := productServiceClient.GetProductName(context.Background(), productRequest)

	if err != nil {
		log.Printf("GetProductName err:%s\n",err.Error())
		return
	}
	productResponse2, err := productServiceClient.GetProductStockStatus(context.Background(), productRequest)
	if err != nil {
		log.Printf("GetProductName err:%s\n",err.Error())
		return
	}
	log.Printf("name:%s,stock_staus:%d",productResponse.ProductName,productResponse2.StockStatus)

}
