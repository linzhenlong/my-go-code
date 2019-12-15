package process

import (
	"errors"
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/common/message"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/server/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 编写一个ServerProcessMsg 函数
// 功能：根据客户端发送的消息种类不同,决定调用哪个函数来处理.
func (processor *Processor)serverProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
	// 登录请求
	case message.LoginMsgType:

		// 创建userProcess 实例
		userProcess := &UserProcess{
			Conn:processor.Conn,
		}
		err = userProcess.ServerProcessLogin(msg)
		break
	case message.RegisterMsgType:
		break
	default:
		return errors.New("消息类型不存在")
	}
	return
}


func (processor *Processor)Handle() (err error){


	//buf := make([]byte, 8096)
	// 读取客户端发送的消息
	for {

		// 将读取消息封装成一个函数.
		// 创建transfer 实例
		transfer := &utils.Transfer{
			Conn:processor.Conn,
		}
		msg, err := transfer.ReadPkg()
		if err != nil {
			return err
		}
		fmt.Println(msg)
		err = processor.serverProcessMsg(&msg)
		if err != nil {
			fmt.Println("serverProcessMsg error=", err)
			return err
		}

	}
}
