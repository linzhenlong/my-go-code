package main

import (
	"context"
	"fmt"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/001-1/client/services"
	"google.golang.org/grpc"
	"log"
)

func main() {
	clientConn, err := grpc.Dial(":6060", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer clientConn.Close()
	request := &services.SearchRequest{
		PageNumber:    2,
		ResultPerPage: 100,
	}
	searchServiceClient := services.NewSearchServiceClient(clientConn)
	ctx := context.Background()
	articles, err := searchServiceClient.GetArticles(ctx, request)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("error_code", articles.ErrorCode, "error_msg", articles.ErrorMsg)
	fmt.Println(articles.Data.Rows)
}
