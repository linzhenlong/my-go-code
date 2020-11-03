package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"log"

	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/001/client/services"
	"google.golang.org/grpc"
)

func main() {

	// 加入证书
	clientTLSFromFile, err := credentials.NewClientTLSFromFile("./keys/server.crt", "linzl.com")
	if err != nil {
		log.Fatalln(err)
	}
	// 带有证书的
	rpcCli, err := grpc.Dial(":6060", grpc.WithTransportCredentials(clientTLSFromFile))

	if err != nil {
		log.Fatalf("err:%v", err)
	}
	defer rpcCli.Close()
	prodClient := services.NewProdServiceClient(rpcCli)

	proRequest := services.ProductRequest{
		ProdId: 20,
	}
	// 调用服务端方法
	proResp, err := prodClient.GetProdStock(context.Background(), &proRequest)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("resp:%#v\n", proResp)

	// 获取商品名称
	proName ,err := prodClient.GetProductName(context.Background(),&proRequest)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("商品名称:%s\n",proName.ProductName)
	//https://www.bilibili.com/video/BV1Fa4y1i7C6?p=4
}
