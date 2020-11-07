package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/006-枚举类型/server/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)

	// 启动grpc服务
	go func() {
		// 创建 grpc server
		grpcServer := grpc.NewServer()

		//注册grpc 服务
		services.RegisterProductServiceServer(grpcServer, new(services.ProductService))

		// 监听6061端口
		listener, err := net.Listen("tcp", ":6061")
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("grpc 启动...")
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	// grpc-gateway http server
	go func() {
		serveMux := runtime.NewServeMux()
		options := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		// 注册服务
		ctx := context.Background()
		endPoint := "localhost:6061"
		// 注册grpc 服务
		err := services.RegisterProductServiceHandlerFromEndpoint(ctx, serveMux, endPoint, options)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 启动http 服务
		httpServer := http.Server{
			Addr:    ":6060",
			Handler: serveMux,
		}
		fmt.Println("启动 grpc-gateway http 服务...")
		httpServer.ListenAndServe()
	}()

	signal.Notify(c, os.Interrupt,os.Kill)
	s := <-c
	fmt.Println("退出了",s)
}

https://www.bilibili.com/video/BV1Fa4y1i7C6?p=10