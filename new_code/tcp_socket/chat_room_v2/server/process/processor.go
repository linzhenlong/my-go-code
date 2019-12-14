package process

import (
	"errors"
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/common/message"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/server/utils"
	"net"
)

// 编写一个ServerProcessMsg 函数
// 功能：根据客户端发送的消息种类不同,决定调用哪个函数来处理.

func serverProcessMsg(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {
	// 登录请求
	case message.LoginMsgType:
		err = ServerProcessLogin(conn, msg)
		break
	case message.RegisterMsgType:
		break
	default:
		return errors.New("消息类型不存在")
	}
	return
}


func Process(conn net.Conn) {

	defer conn.Close()
	//buf := make([]byte, 8096)
	// 读取客户端发送的消息
	for {

		// 将读取消息封装成一个函数.
		msg, err := utils.ReadPkg(conn)
		if err != nil {
			return
		}
		fmt.Println(msg)
		err = serverProcessMsg(conn, &msg)
		if err != nil {
			fmt.Println("serverProcessMsg error=", err)
			return
		}

	}
}
