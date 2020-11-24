package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/008日期类型-订单模型/server/services"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)

	//  启动grpc 服务
	go func() {
		// 创建server
		grpcServer := grpc.NewServer()

		// 注册商品服务
		services.RegisterProductServiceServer(grpcServer, new(services.ProductService))

		// 注册订单服务
		services.RegisterOrderServiceServer(grpcServer, new(services.OrderService))

		// 监听端口
		listener, err := net.Listen("tcp", ":6061")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("grpc 启动啦....")
		err = grpcServer.Serve(listener)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	go func() {
		// 创建grpc-gateway 路由
		serveMux := runtime.NewServeMux()

		// 注册服务
		ctx := context.Background()
		options := []grpc.DialOption{
			grpc.WithInsecure(),
		}

		// 注册商品服务
		err := services.RegisterProductServiceHandlerFromEndpoint(ctx, serveMux, "localhost:6061", options)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 注册订单服务
		err = services.RegisterOrderServiceHandlerFromEndpoint(ctx, serveMux, "localhost:6061", options)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("grpc gateway 启动http server")
		// http server
		httpServer := http.Server{
			Addr:    ":6060",
			Handler: serveMux,
		}
		err = httpServer.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("退出了", s)
}

//
