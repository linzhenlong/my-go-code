package services

import (
	"context"
	"errors"
)

// UserService ...
type UserService struct {

}
// GetUserName 获取用户名
func(u *UserService) GetUserName(ctx context.Context, ur *UserRequest) (*UserResponse, error) {
	// 业务逻辑
	var userName string
	if ur.UserId == 18 {
		userName = "张三"
	} else {
		userName = "王五"
	}
	resp := &UserResponse{
		User: &UserInfo{
			UserName: userName,
		},
	}
	return resp, nil
}
// GetUserInfo
func (u *UserService)GetUserInfo(ctx context.Context, ur *UserRequest) (*UserInfo, error) {
	userInfo := &UserInfo{}
	if ur.UserId == 18 {
		userInfo.UserId = ur.UserId
		userInfo.UserName = "张三"
		userInfo.Age = 20
	} else {
		userInfo.UserId = ur.UserId
		userInfo.UserName = "李四"
		userInfo.Age = 18
	}
	if ur.UserId > 100 {
		return userInfo, errors.New("模拟一个err")
	}
	return userInfo, nil
}
// GetUserID ...
func (u *UserService)GetUserID(ctx context.Context, ur *UserRequestByName) (*UserID, error) {
	userId := 0
	if ur.UserName == "张三" {
		userId = 18
	}
	return &UserID{
		UserId: int32(userId),
	},nil
}