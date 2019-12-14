package main

import (
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/tcp_socket/chat_room_v2/server/process"
	"net"
)

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

		go process.Process(conn)
	}
}
