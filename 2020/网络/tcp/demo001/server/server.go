package main

import "net"

import "fmt"

import "strings"


// 实现一下接口.
type myClient struct {
	
}
func (m *myClient)String() string{
	return "客户端ip"
}
func(m *myClient)Network()string {
	return "客户端network"
}

func main() {
	// 创建tcp服务端监听
	listener, err := net.Listen("tcp",":6060")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	// 2.服务端不断的等待处理请求
	for {

		// 阻塞等待
		conn,err := listener.Accept()
		if err !=nil {
			fmt.Println(err)
			continue
		}
		go ClientConn(conn)

	}
}
// ClientConn ...
func ClientConn(conn net.Conn) {
	defer conn.Close()

	// 客户端地址
	 myClient := myClient{}
	//var clientIp net.Addr
	clientIP := conn.RemoteAddr()
	clientIP = &myClient
	fmt.Println(clientIP.String(),"链接成功")
	
	fmt.Println(clientIP.Network(),"链接了")

	// 缓冲区
	buf := make([]byte,1024)
	for {
		// n 读取的长度
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 读取出内容
		result := string(buf[:n])
		fmt.Printf("我接收到了数据:%s\n",result)

		// 接收到exit 退出链接
		if result == "exit" {
			fmt.Println(clientIP,"退出链接")
			return
		}
		
		// 回复客户端
		conn.Write([]byte(strings.ToUpper(result)))
	}
	
	

}