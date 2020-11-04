package main

import (
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/001-1/server/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	server := grpc.NewServer()
	// 注册服务
	searchService := &services.SearchService{}
	services.RegisterSearchServiceServer(server, searchService)

	// 拉起服务监听6060端口
	listener, err := net.Listen("tcp", ":6060")
	if err != nil {
		log.Fatal(err)

	}
	server.Serve(listener)

	/*serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request)
		server.ServeHTTP(writer, request)
	})
	httpServer := http.Server{
		Addr:    ":6060",
		Handler: serveMux,
	}
	httpServer.ListenAndServe()*/

}
