package main

import (
	"net"

	"fmt"

	"github.com/linzhenlong/my-go-code/2020/mysql-middleware/util"
	"github.com/siddontang/go-mysql/server"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:3309")
	if err != nil {
		fmt.Println("Listen err", err)
		return
	}
	for {
		c, _ := listen.Accept()
		go func() {
			conn, err := server.NewConn(c, "root", "123456", util.NewMyHandler())
			if err != nil {
				fmt.Println("NewConn err", err)
				return
			}
			for {
				conn.HandleCommand()
			}
		}()
	}
}
