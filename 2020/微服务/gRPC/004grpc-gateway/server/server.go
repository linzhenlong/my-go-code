package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/004grpc-gateway/server/helper"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/004grpc-gateway/server/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// 通过两个携程 分别拉起 grpc-server 及grpc-gateway http server

	exitChan := make(chan os.Signal)
	go grpcServer()
	go grpcGateWayHttpServer()
	signal.Notify(exitChan,os.Interrupt)
	s := <- exitChan
	log.Println("exit",s)
}

func grpcServer() {
	// ca验证
	grpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCred()))
	// 注册服务
	productService := &services.ProductService{}
	services.RegisterProductServiceServer(grpcServer,productService)

	searchService := &services.SearchService{}
	services.RegisterSearchServiceServer(grpcServer,searchService)

	listener, err := net.Listen("tcp", ":6061")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("grpcServer running")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
	/*serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		grpcServer.ServeHTTP(writer, request)
	})
	httpServer := http.Server{
		Addr:    ":6060",
		Handler: serveMux,
	}
	err := httpServer.ListenAndServeTLS(helper.ServerCrtFile, helper.ServerKeyFile)
	if err != nil {
		log.Println(err)
	}*/
}


func grpcGateWayHttpServer() {
	// grpc-gateway 路由
	serveMux := runtime.NewServeMux()

	// gateway 里有客户端证书访问grpc 服务
	option := []grpc.DialOption{
		grpc.WithTransportCredentials(helper.GetClientCred()),
	}
	// 注册服务
	err := services.RegisterProductServiceHandlerFromEndpoint(
		context.Background(),
		serveMux,
		"localhost:6061",
		option,
	)
	if err != nil {
		log.Fatal(err)
	}
	err2 := services.RegisterSearchServiceHandlerFromEndpoint(
		context.Background(),
		serveMux,
		"localhost:6061",
		option,
	)
	if err2 != nil {
		log.Fatal(err2)
	}
	// 启动HTTP SERVER
	httpServer := http.Server{
		Addr:    ":6060",
		Handler: serveMux,
	}

	log.Println("grpcGateWayHttpServer running")
	err3 := httpServer.ListenAndServe()
	if err3 != nil {
		log.Fatal(err3)
	}
}




//https://www.bilibili.com/video/BV1Fa4y1i7C6?p=7