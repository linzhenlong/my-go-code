package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/005grpc-gateway-商品列表/client/helper"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/005grpc-gateway-商品列表/server/services"
	"google.golang.org/grpc"
	"log"
)

func main() {
	size := flag.Int("size", 10, "输入size")
	flag.Parse()
	clientConn, err := grpc.Dial(":6061", grpc.WithTransportCredentials(helper.GetClientCred()))
	if err != nil {
		log.Println(err)
		return
	}
	defer clientConn.Close()
	// 创建client
	productServiceClient := services.NewProductServiceClient(clientConn)

	// 获取商品列表
	ctx := context.Background()
	listRequest := &services.QuerySize{
		Size: int32(*size),
	}
	list, err := productServiceClient.GetProductProductList(ctx, listRequest)
	if err != nil {
		log.Println(err)
		return
	}
	for _,product := range list.List {
		fmt.Printf("商品名称:%s,库存:%d\n",product.ProductName,product.StockStatus)
	}
}
