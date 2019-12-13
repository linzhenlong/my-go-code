package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)



func Process(conn net.Conn, msgMap map[string]string)  {
	// 关闭链接
	defer conn.Close()

	// 循环接收客户端发送的数据
	for {
		fmt.Printf("服务器端在响应客户端【%s】的输入\n",conn.RemoteAddr().String())
		// 创建一个byte切片
		buf := make([]byte,1024)

		// 1.等待客户端通过tcp链接发生信息给服务端
		// 2.如果客户端没有发送信息,那么协程就一直阻塞在这里
		n , err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已经断开连接")
			} else {
				fmt.Printf("conn.Read error:%v\n", err)
			}
			return
		}
		// 显示客户端发送的信息
		// buf[:n] 表示显示客户端实际发送的内容，因为在make 切片时的长度是1024
		clientMsg :=  string(buf[:n])
		clientMsg = strings.Trim(clientMsg,"\r\n")
		fmt.Printf("客户端发送了【%s】\n",clientMsg)

		fmt.Println(msgMap[clientMsg])

		// 判断map中是否含有key.
		if val,ok := msgMap[clientMsg];ok {
			_,err :=conn.Write([]byte(val))
			if err!=nil {
				fmt.Println("conn.Write error",err)
			}
		} else {
			_,err :=conn.Write([]byte("你说啥？？"))
			if err!=nil {
				fmt.Println("conn.Write error",err)
			}
		}



	}
}

func main()  {

	fmt.Println("服务端开始监听,0.0.0.0:8888....")

	// 1.tcp 标识网络协议是tcp
	// 2.0.0.0.0:8888 本地监听8888端口
	listen, err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Printf("Listen err:%v\n", err)
		return
	}

	defer func() {
		err = listen.Close() // 关闭
		if err != nil {
			fmt.Printf("listen.Close() error:%v", err)
		}
	}()

	msgMap := make(map[string]string, 5)
	msgMap["你是谁?"] = "我是小冰"
	msgMap["你的性别是"] = "女"

	// 循环等待客户端链接
	for {
		// 等待客户端链接
		fmt.Println("等待客户端链接。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("listen.Accept() error:%v\n", err)
			continue
		}

		fmt.Printf("listen.Accept() success,conn=:%v\n",conn)
		fmt.Printf("客户端ip=%v\n",conn.RemoteAddr().String())

		// 起协程处理
		go Process(conn,msgMap)

	}

}