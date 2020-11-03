package main

import (
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/001/services"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// grpc 创建server

	rpcServer := grpc.NewServer()

	// 注册服务
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	// 监听6060端口
	listener, _ := net.Listen("tcp", ":6060")

	rpcServer.Serve(listener)

}
