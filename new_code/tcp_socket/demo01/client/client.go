package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err !=nil {
		fmt.Printf("net.Dial()链接失败,error=:[%v]\n", err)
		return
	}
	fmt.Println(conn)

	// 客户端从键盘输入一句话给服务端
	// os.Stdin 代表标准输入,可以理解为从终端输入
	reader := bufio.NewReader(os.Stdin)

	// 从终端读取用户输入,并发送给服务器
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("ReadString读取失败error=%v\n", err)
			return
		}
		if strings.Trim(msg,"\r\n") == "exit"{
			break
		}
		// 在将msg 发给服务端
		n, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("conn.Write error=%v\n", err)
		}
		fmt.Printf("成功写入%d个byte\n", n)

		serverMsg := make([]byte, 1024)
		serverN,err := conn.Read(serverMsg)
		if err !=nil {
			fmt.Printf("conn.Readerror=%v\n", err)
			continue
		}
		fmt.Printf("[服务器端返回===>%s]\n",string(serverMsg[:serverN]))

	}


}
