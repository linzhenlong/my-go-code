package main

import (
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/client/process"
	"os"
)

var (
	Key int // 用户输入选择
	IsExit bool // 是否退出

	UserId int
	UserPassword string
)
func main()  {
	var loop bool
	loop = true
	for loop {
		fmt.Println("---------欢迎登录多人聊天系统--------")
		fmt.Println("-------------1:登录聊天室--------")
		fmt.Println("-------------2:注册用户--------")
		fmt.Println("-------------3:退出--------")
		fmt.Println("请选择（1-3）：")

		fmt.Scanf("%d\n", &Key)
		switch Key {
		case 1:
			fmt.Println("登录聊天室~~~")
			fmt.Println("用户id:")
			fmt.Scanf("%d\n", &UserId)
			fmt.Println("用户密码:")
			fmt.Scanf("%s\n", &UserPassword)


			userProcess := &process.UserProcess{}
			_ = userProcess.Login(UserId, UserPassword)

			//loop = false
		case 2:
			fmt.Println("用户注册")
			//loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误,请选择（1-3）")
		}
	}
}


