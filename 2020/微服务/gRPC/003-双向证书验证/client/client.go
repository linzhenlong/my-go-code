package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/003-双向证书验证/server/services"
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
		log.Println(err)
	}
	defer clientConn.Close()

	// 创建grpc 客户端
	userServiceClient := services.NewUserServiceClient(clientConn)

	// 调用方法
	ctx := context.Background()
	requestByName := &services.UserRequestByName{
		UserName: "张三",
	}
	userIDResp, err := userServiceClient.GetUserID(ctx, requestByName)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("userServiceClient.GetUserID", userIDResp)
	userRequest := &services.UserRequest{
		UserId: 18,
	}
	getUserInfo, err := userServiceClient.GetUserInfo(ctx, userRequest)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("GetUserInfo",getUserInfo)

	userRequest2 := &services.UserRequest{
		UserId: 10000,
	}
	getUserInfo2, err := userServiceClient.GetUserInfo(ctx, userRequest2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("GetUserInfo2",getUserInfo2)

	getUserName, err := userServiceClient.GetUserName(ctx, &services.UserRequest{
		UserId: 18,
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("GetUserName",getUserName.User.UserName)
}

// https://www.bilibili.com/video/BV1Fa4y1i7C6?p=7
