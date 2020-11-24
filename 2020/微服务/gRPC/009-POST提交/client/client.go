package main

import (
	"context"
	"fmt"
	"github.com/linzhenlong/my-go-code/2020/微服务/gRPC/009-POST提交/server/services"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	clientConn, err := grpc.Dial(":6061", grpc.WithInsecure())
	if err != nil {
		log.Printf("%v\n",err)
		return
	}
	defer clientConn.Close()

	serviceClient := services.NewUserServiceClient(clientConn)
	ctx := context.Background()
	users := make([]*services.UserInfo,0)
	for i:=0;i<5;i++ {
		user := services.UserInfo{
			UserId: int32(i),
		}
		users = append(users,&user)
	}
	request := services.UserScoreRequest{
		Users: users,
	}
	resp, err := serviceClient.GetUserScore(ctx, &request)
	if err != nil {
		log.Printf("%v\n",err)
		return
	}
	fmt.Println(resp)


	fmt.Println("服务端流模式调用")
	// 客户端请求服务端流模式方法
	serverStream, err := serviceClient.GetUserScoreByServerStream(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	// 循环读
	for {
		userScoreResponse, err := serverStream.Recv()
		// 结束了
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(userScoreResponse)
	}

	fmt.Println("客户端流模式调用=======")
	// 客户端流模式调用
	clientStreamClient, err := serviceClient.GetUserScoreByClientStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 循环5次每次发5条 批量
	for j:=1;j<=5;j++{
		request2 := &services.UserScoreRequest{
			Users: users,
		}
		users2 := make([]*services.UserInfo,0)
		for i:=0;i<5;i++ {
			user := services.UserInfo{
				UserId: int32(i),
			}
			users2 = append(users2,&user)
		}
		err := clientStreamClient.Send(request2)
		if err != nil {
			log.Println(err)
		}
	}
	// 最后获取响应
	recv, err := clientStreamClient.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(recv)

	fmt.Println("双向流模式")
	// 双向流模式
	scoreByStream, err := serviceClient.GetUserScoreByStream(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// 循环5次每次发5条 批量
	for j:=1;j<=5;j++{
		request2 := &services.UserScoreRequest{
			Users: users,
		}
		users2 := make([]*services.UserInfo,0)
		for i:=0;i<5;i++ {
			user := services.UserInfo{
				UserId: int32(i),
			}
			users2 = append(users2,&user)
		}
		err := scoreByStream.Send(request2)
		if err != nil {
			log.Println(err)
		}
		scoreResponse, err := scoreByStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(scoreResponse)
	}

}
