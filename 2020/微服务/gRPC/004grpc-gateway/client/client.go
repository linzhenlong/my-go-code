package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/004grpc-gateway/client/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func main() {

	// ca 验证
	var (
		crtFile = "keys/client.pem"
		keyFile = "keys/client.key"
		caPem = "keys/ca.pem"
	)
	keyPair, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Println(err)
	}
	// 创建证书池
	certPool := x509.NewCertPool()
	caFile, _ := ioutil.ReadFile(caPem)
	certPool.AppendCertsFromPEM(caFile)
	transportCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{keyPair}, // 客户端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	clientConn, err := grpc.Dial(":6060",grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		log.Fatal(err)
	}
	defer clientConn.Close()

	// 初始化客户端
	productServiceClient := services.NewProductServiceClient(clientConn)

	searchServiceClient := services.NewSearchServiceClient(clientConn)



	// 获取库存
	ctx := context.Background()
	productRequest := &services.ProductRequest{}
	stockStatus, err := productServiceClient.GetProductStockStatus(ctx, productRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("库存是",stockStatus.StockStatus)

	productName, err := productServiceClient.GetProductName(ctx, productRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("商品名称是",productName.ProductName)

	// 文章列表
	articles, err := searchServiceClient.GetArticles(ctx, &services.SearchRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("articles",articles)

}
