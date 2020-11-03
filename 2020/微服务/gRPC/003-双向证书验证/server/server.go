package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/003-双向证书验证/server/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// ca验证
	var (
		crtFile = "keys/server.pem"
		keyFile = "keys/server.key"
		caPem = "keys/ca.pem"
	)
	keyPair, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	caFile, err := ioutil.ReadFile(caPem)
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(caFile)
	newTLS := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{keyPair}, // 服务端证书
		ClientAuth:   tls.RequireAnyClientCert,  // 客户端验证需要证书，双向验证
		ClientCAs:    certPool, // 证书池
	})

	// 启动grpc server
	grpcServer := grpc.NewServer(grpc.Creds(newTLS))

	// 注册服务
	userService := &services.UserService{}

	// 注册服务
	services.RegisterUserServiceServer(grpcServer, userService)

	// 创建路由
	httpServeMux := http.NewServeMux()
	
	httpServeMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request)
		grpcServer.ServeHTTP(writer, request)
	})

	// 启动http服务
	httpServer := http.Server{
		Addr:    ":6060",
		Handler: httpServeMux,
	}
	httpServer.ListenAndServeTLS(crtFile, keyFile)

}
