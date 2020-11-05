package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/005grpc-gateway-商品列表/server/helper"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/005grpc-gateway-商品列表/server/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	// 起两个携程拉取服务
	c := make(chan os.Signal)

	go func() {
		// 启动grpc server
		grpcServer := grpc.NewServer(grpc.Creds(helper.GetClientCred()))

		// 注册服务
		services.RegisterProductServiceServer(grpcServer,new(services.ProductService))
		services.RegisterSearchServiceServer(grpcServer,new(services.SearchService))

		// 拉起服务
		listener, err := net.Listen("tcp", ":6061")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("拉起grpc server success")
		grpcServer.Serve(listener)
	}()

	go func() {
		// 启动grpc-gateway http server
		serveMux := runtime.NewServeMux()

		options := []grpc.DialOption{
			grpc.WithTransportCredentials(helper.GetClientCred()),
		}
		// 注册服务
		ctx := context.Background()
		endPoint := "localhost:6061"
		err := services.RegisterSearchServiceHandlerFromEndpoint(ctx, serveMux, endPoint, options)
		if err != nil {
			log.Println("注册search services fail",err)
		}
		err = services.RegisterProductServiceHandlerFromEndpoint(ctx, serveMux, endPoint, options)
		if err != nil {
			log.Println("注册Product services fail",err)
		}

		// 拉起http server
		htppServer := http.Server{
			Addr:    ":6060", // 监听6060端口
			Handler: serveMux,
		}
		log.Println("拉起grpc-gateway http-server success")
		err = htppServer.ListenAndServe()
		if err != nil {
			log.Println("ListenAndServe err",err)
		}
	}()
	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println("算是优雅退出么？",s)
}
