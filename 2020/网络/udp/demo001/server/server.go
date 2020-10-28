package main

import "net"

import "fmt"

func main() {
	UDPconn, err := net.ListenUDP("udp",&net.UDPAddr{
		//IP: []byte(`127.0.0.1`),
		IP: net.IPv4(127,0,0,1),
		Port: 6060,
	})
	if err != nil {
		panic(err)
	}
	defer UDPconn.Close()
	fmt.Println(UDPconn.LocalAddr().String(),"启动了...")
	for {
		// 缓冲区
		var data [1024]byte
		n, UDPAddr,err := UDPconn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("客户端NetWork:",UDPAddr.Network())
		fmt.Println("客户端Addr:",UDPAddr.String())
		fmt.Printf("客户端发送了:%s\n",string(data[0:n]))
		UDPconn.WriteToUDP([]byte(`6666`),UDPAddr)

	}
}

