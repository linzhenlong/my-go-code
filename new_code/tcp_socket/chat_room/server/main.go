package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room/common/message"
	"github.com/pkg/errors"
	_ "io"
	"net"
)

func readPkg(conn net.Conn) (message.LoginResMsg, error) {
	buf := make([]byte, 8096)
	n, err := conn.Read(buf[0:4])
	response := message.LoginResMsg{
		ErrorCode: 500,
		ErrorMsg:  "",
	}
	if err != nil {
		response.ErrorMsg = err.Error()
		response.ErrorCode = 500
		return response, errors.New("conn.Read head出错")
	}

	// 根据读到的buf[0:4] 转成uint32类型的
	pkgLen := binary.BigEndian.Uint32(buf[0:4])

	// 根据pkgLen 读取消息内容
	// buf[0:pkgLen] 表示获取客户端发送的数据包读取出pkgLen长度的数据，放到buf里面去.
	n, err = conn.Read(buf[0:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		response.ErrorMsg = err.Error()
		response.ErrorCode = 500
		return response, errors.New("conn.Read 数据包出错")
	}

	// 将buf[0:pkgLen] 反序列化
	requestMsg := message.Message{}
	err = json.Unmarshal(buf[0:pkgLen], &requestMsg)
	if err != nil {
		response.ErrorMsg = err.Error()
		response.ErrorCode = 500
		return response, errors.New("反序列化buf[0:pkgLen]出错")
	}
	switch requestMsg.Type {
	case message.LoginMsgType:
		loginMsg := message.LoginMsg{}
		err = json.Unmarshal([]byte(requestMsg.Data), &loginMsg)
		if err != nil {
			response.ErrorMsg = err.Error()
			response.ErrorCode = 500
			return response, errors.New("反序列化requestMsg.Data出错")
		}
		if loginMsg.UserId == 12345 && loginMsg.UserPwd == "abc" {
			response.ErrorCode = 0
			response.ErrorMsg = "success"
			return response, nil
		} else {
			response.ErrorCode = 1
			response.ErrorMsg = "用户名或密码错误"
			return response, nil
		}
	default:
		response.ErrorMsg = "未知类型"
		response.ErrorCode = 500
		// 自定义错误

		return response, errors.New("类型有误")
	}
}

func process(conn net.Conn, ) {

	defer conn.Close()
	//buf := make([]byte, 8096)
	// 读取客户端发送的消息
	for {

		// 将读取消息封装成一个函数.
		msg, err := readPkg(conn)
		if err != nil {
			return
		}
		fmt.Println(msg)
		//n, err := conn.Read(buf[0:4])
		//if n !=4 || err != nil {
		//	if err == io.EOF {
		//		fmt.Println("客户端断开链接")
		//		return
		//	} else {
		//		fmt.Println("conn.Read error", err)
		//		continue
		//	}
		//}
		//fmt.Println("读到的buf=",string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器在8889端口监听.....")

	listener, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	// 一旦监听成功，等待客户端链接服务器
	for {
		fmt.Println("等待客户端链接.....")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept() error=", err)
			continue
		}

		fmt.Println("客户端连接成功，客户端ip=", conn.RemoteAddr())

		// 链接成功，启协程和客户端保持数据通讯

		go process(conn)
	}
}
