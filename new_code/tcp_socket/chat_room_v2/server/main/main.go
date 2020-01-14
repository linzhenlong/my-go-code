package main

import (
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/server/model"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/server/process"
	"net"
	"time"
)

func handle(conn net.Conn)  {
	defer conn.Close()

	// 实例化processor
	processor := &process.Processor{
		Conn:conn,
	}
	err := processor.Process()
	if err != nil {
		fmt.Println("processor.Process() error=", err)
		return
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

// 初始化redis链接
func init()  {

	// 初始化redis.
	initPool("127.0.0.1:6379",8,10,time.Second*100)

	// 初始化userDAO
	initUserDao()
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

		go handle(conn)
	}

	// 该看视频 331_尚硅谷_Go核心编程_海量用户通讯系统-显示在线用户列表(2)
}
