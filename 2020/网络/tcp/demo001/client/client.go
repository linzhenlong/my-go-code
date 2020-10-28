package main

import "net"

import "fmt"

func main() {
	// 创建链接
	conn,err := net.Dial("tcp",":6060")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	
	// 带缓冲的
	buf := make([]byte, 1024)
	for {
		fmt.Println("请输入发送的内容：")
		fmt.Scan(&buf)
		fmt.Printf("发送的内容是:%s\n",string(buf))
		// 发送数据
		conn.Write(buf)
		// 接收服务响应
		res := make([]byte,1024)
		n ,err := conn.Read(res)
		if err !=nil {
			fmt.Println(err)
			return
		}
		resp := string(res[:n])
		fmt.Println("服务端说:",resp)
	}
}