package main

import "net/rpc"

import "log"

import "fmt"

/**
rpc 客户端
**/

// Params ...
// 2.声明参数的结构体
type Params struct {
	Width  int // 长
	Hegiht int // 宽
}

func main() {
	// 1.连接远程rpc 服务
	rpcClient, err := rpc.DialHTTP("tcp", "127.0.0.1:6060")
	if err != nil {
		log.Println(err)
		return
	}
	defer rpcClient.Close()

	// 2.调用远程的方法
	params := Params{
		Width:  2,
		Hegiht: 2,
	}
	area := 0
	err = rpcClient.Call("Rect.Area", params, &area)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("area:", area)

	zhouchang := 0
	rpcClient.Call("Rect.ZhouChang", params, &zhouchang)
	fmt.Println("ZhouChang:", zhouchang)
}
