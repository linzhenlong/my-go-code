package main

import (
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/001/server/services"
	"google.golang.org/grpc"

)

func main() {

	// 添加证书
	crtFile, err := credentials.NewServerTLSFromFile("./keys/server.crt", "./keys/server_no_password.key")
	if err != nil {
		log.Printf("credentials NewServerTLSFromFile err:%v",err)
	}

	// 构建grpc server
	grpcSrv := grpc.NewServer(grpc.Creds(crtFile))

	// 注册服务

	prodServices := &services.ProdService{

	}
	services.RegisterProdServiceServer(grpcSrv, prodServices)

	// 监听6060端口
	listener, err := net.Listen("tcp", ":6060")
	if err != nil {
		log.Fatalf("net.Listen 6060 err:%v\n", err)
	}
	grpcSrv.Serve(listener)


}
