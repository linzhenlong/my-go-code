package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/009-POST提交/server/services"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)

	go func() {

		grpcServer := grpc.NewServer()
		// 注册服务
		services.RegisterOrderServiceServer(grpcServer,new(services.OrderService))
		services.RegisterProductServiceServer(grpcServer,new(services.ProductService))
		services.RegisterUserServiceServer(grpcServer,new(services.UserService))

		// 监听端口
		listener, err := net.Listen("tcp", ":6061")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("拉取grpc服务")
		grpcServer.Serve(listener)
	}()

	go func() {
		serveMux := runtime.NewServeMux()
		//注册服务
		ctx := context.Background()
		endPoint := "localhost:6061"
		options :=[]grpc.DialOption{
			grpc.WithInsecure(),
		}
		err := services.RegisterOrderServiceHandlerFromEndpoint(ctx, serveMux, endPoint, options)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = services.RegisterProductServiceHandlerFromEndpoint(ctx, serveMux, endPoint, options)
		if err != nil {
			fmt.Println(err)
			return
		}
		httpServer := http.Server{
			Addr:    ":6060",
			Handler: serveMux,
		}
		fmt.Println("拉起grpc-gateway http server")
		httpServer.ListenAndServe()
	}()
	signal.Notify(c, os.Interrupt, os.Kill)
	fmt.Println("退出了",<-c)
}
