package main

import (
	"context"
	"log"

	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/001/client/services"
	"google.golang.org/grpc"
)

func main() {
	rpcCli, err := grpc.Dial(":6060", grpc.WithInsecure())

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
	//https://www.bilibili.com/video/BV1Fa4y1i7C6?p=4
}
