package services

import (
	"context"
	"io"
)

type UserService struct {

}

func (u *UserService) GetUserScore(ctx context.Context, request *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo,0)
	for _,user := range request.Users{
		user.UserScore = score
		score++
		users = append(users,user)
	}
	return &UserScoreResponse{Users: users},nil
}

// GetUserScoreByServerStream 服务端流模式方法.
func (u *UserService) GetUserScoreByServerStream(request *UserScoreRequest, stream UserService_GetUserScoreByServerStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo,0)
	for index,user := range request.Users{
		user.UserScore = score
		score++
		users = append(users,user)
		// 分批每两次发一条
		if (index+1) %2 == 0 && index >0 {
			err := stream.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			// 每次清空切片
			users = (users)[0:0]
		}
	}
	// users 里面还有值在发一次
	if len(users) > 0 {
		err := stream.Send(&UserScoreResponse{Users: users})
		if err != nil {
			return err
		}
	}
	return nil
}

// GetUserScoreByClientStream 客户端流模式方法.
func (u *UserService)GetUserScoreByClientStream(stream UserService_GetUserScoreByClientStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo,0)
	for {
		request, err := stream.Recv()

		// 接收完毕,发送出去
		if err == io.EOF {
			return stream.SendAndClose(&UserScoreResponse{Users: users})
		}
		if err != nil {
			return err
		}

		// 服务端的业务逻辑
		for _,user := range request.Users{
			user.UserScore = score
			score++
			users = append(users,user)
		}
	}

	return nil
}
// GetUserScoreByStream 双向流
func (u *UserService)GetUserScoreByStream(stream UserService_GetUserScoreByStreamServer) error {

	var score int32 = 101
	users := make([]*UserInfo,0)
	for {
		request, err := stream.Recv()

		// 接收完毕,发送出去
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// 服务端的业务逻辑
		for _,user := range request.Users{
			user.UserScore = score
			score++
			users = append(users,user)
		}
		err = stream.Send(&UserScoreResponse{Users: users})
		if err != nil {
			return err
		}
		users = (users)[0:0]
	}
}