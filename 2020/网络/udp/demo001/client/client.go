package main

import (
	"fmt"
	"net"
)

func main() {
	UDPConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 6060,
	})
	if err != nil {
		panic(err)
	}
	defer UDPConn.Close()

	// 写数据到服务端
	_, err = UDPConn.Write([]byte("老铁"))
	if err != nil {
		panic(err)
	}

	// 接收数据
	data := make([]byte, 16)
	n, addr, err := UDPConn.ReadFromUDP(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(n,addr.Network(),addr.String())
	fmt.Println("data:",string(data))

}
