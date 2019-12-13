package view

import (
	"fmt"
	"github.com/linzhenlong/mygo/go_dev/new_code/tcp_socket/chat_room/client/service"
)

type UserView struct {
	Key int // 用户输入选择
	IsExit bool // 是否退出

	UserId int
	UserPassword string

	userService *service.UserService


}

func (userView *UserView)login() {
	fmt.Println("登录聊天室~~~")
	fmt.Println("用户id:")
	fmt.Scanf("%d\n",&userView.UserId)
	fmt.Println("用户密码:")
	fmt.Scanf("%s\n",&userView.UserPassword)

	userView.userService = service.NewUserService(userView.UserId,userView.UserPassword)
	err := userView.userService.Login()
	if err!= nil {
		panic(err)
	}

}

func (userView *UserView)register()  {
	fmt.Println("注册新用户")

}


func (userView *UserView)ShowMainView()  {

	var loop bool
	loop = true
	for loop {
		fmt.Println("---------欢迎登录多人聊天系统--------")
		fmt.Println("-------------1:登录聊天室--------")
		fmt.Println("-------------2:注册用户--------")
		fmt.Println("-------------3:退出--------")
		fmt.Println("请选择（1-3）：")

		fmt.Scanf("%d\n", &userView.Key)
		switch userView.Key {
		case 1:
			userView.login()
			loop = false
		case 2:
			userView.register()
			loop = false
		case 3:
			userView.IsExit = true
		default:
			fmt.Println("输入有误,请选择（1-3）")
		}

		if userView.IsExit {
			break;
		}


	}

}