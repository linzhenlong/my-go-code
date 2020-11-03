package main

import (
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/002/server/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
)

func main() {

	var (
		crtFile = "keys/server.crt"
		keyFile = "keys/server_no_password.key"
	)
	serverTLSFromFile, err := credentials.NewServerTLSFromFile(crtFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	// 创建grpc 服务
	server := grpc.NewServer(grpc.Creds(serverTLSFromFile))

	productService := &services.ProductService{}

	// 注册服务
	services.RegisterProductServiceServer(server,productService)

	/*listener, err := net.Listen("tcp", ":6060")
	if err != nil {
		log.Fatal(err)
	}*/

	// 创建路由
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request)
		server.ServeHTTP(writer, request)
	})

	// 启动http server
	httpServer := http.Server{
		Addr:    ":6060",
		Handler: serveMux,
	}
	httpServer.ListenAndServeTLS(crtFile, keyFile)
}
