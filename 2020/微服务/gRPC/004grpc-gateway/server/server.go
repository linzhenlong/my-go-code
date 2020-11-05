package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/004grpc-gateway/server/services"
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
	grpcServer := grpc.NewServer(grpc.Creds(newTLS))

	// 注册服务
	productService := &services.ProductService{}
	services.RegisterProductServiceServer(grpcServer,productService)

	searchService := &services.SearchService{}
	services.RegisterSearchServiceServer(grpcServer,searchService)

	/*listener, err := net.Listen("tcp", ":6060")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listener)*/
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		grpcServer.ServeHTTP(writer, request)
	})
	httpServer := http.Server{
		Addr:    ":6060",
		Handler: serveMux,
	}
	err = httpServer.ListenAndServeTLS(crtFile, keyFile)
	if err != nil {
		log.Println(err)
	}

}


//https://www.bilibili.com/video/BV1Fa4y1i7C6?p=7