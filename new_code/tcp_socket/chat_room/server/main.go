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


func readPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 4096)
	fmt.Println("读取客户端发送消息")

	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}

	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		return
	}
	return
}

// 编写一个函数serverProcessLogin函数,专门处理登录请求
func serverProcessLogin(conn net.Conn, msg *message.Message) (err error) {

	// 核心代码
	// 1. 先从msg中取出msg.Data,并直接反序列化成LoginMsg
	loginMsg := message.LoginMsg{}

	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		return
	}

	// 声明一个responseMsg
	var responseMsg message.Message
	responseMsg.Type = message.LoginResMsgType

	// 再声明一个 LongResMsg
	loginResMsg := message.LoginResMsg{}

	// 如果用户id 为100,密码=123456 ,认为合法否则不合法
	if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
		// 登录成功返回200状态码
		loginResMsg.ErrorCode = 200
		loginResMsg.ErrorMsg = "success"
	} else {
		// 登录错误返回500错误码
		loginResMsg.ErrorCode = 500
		loginResMsg.ErrorMsg = "用户不存在"
	}

	// 序列化一下
	loginResMsgJson, err := json.Marshal(loginResMsg)

	if err != nil {
		return
	}

	// 赋值给responseMsg
	responseMsg.Data = string(loginResMsgJson)

	// 对responseMsg 序列化
	responseMsgJson, err := json.Marshal(responseMsg)
	if err != nil {
		return
	}

	// 发送responseMsgJson数据，将他封装成到一个writePkg()函数
	err = writePkg(conn, responseMsgJson)
	return
}

// 发送数据包函数.
func writePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	// 将一个int 转成字节切片
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf[0:4]) error", err)
		return
	}

	// 发送数据本身
	writeDataLen, err := conn.Write(data)
	if uint32(writeDataLen) != pkgLen || err != nil {
		fmt.Println("conn.Write(data) error", err)
		return
	}
	return

}

// 编写一个ServerProcessMsg 函数
// 功能：根据客户端发送的消息种类不同,决定调用哪个函数来处理.

func serverProcessMsg(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {
	// 登录请求
	case message.LoginMsgType:
		err = serverProcessLogin(conn, msg)
		break
	case message.RegisterMsgType:
		break
	default:
		return errors.New("消息类型不存在")
	}
	return
}

func process(conn net.Conn) {

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
		err = serverProcessMsg(conn, &msg)
		if err != nil {
			fmt.Println("serverProcessMsg error=",err)
			return
		}
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
