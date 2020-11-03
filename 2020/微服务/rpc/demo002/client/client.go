package main

import "net/rpc"

import "log"

import "fmt"

// Params ... 参数
type Params struct {
	X int
	Y int
}
type Resp struct {
	Chengji int
	Shang   float64
	Yushu   int
}

func main() {
	rpcClient, err := rpc.DialHTTP("tcp", "127.0.0.1:6060")
	if err != nil {
		log.Println(err)
		return
	}
	defer rpcClient.Close()
	params := &Params{
		X: 10,
		Y: 0,
	}
	resp := &Resp{}
	var (
		jia  int
		jian int
	)
	rpcClient.Call("Calc.Jia", params, &jia)
	rpcClient.Call("Calc.Jian", params, &jian)
	rpcClient.Call("Calc.Cheng", params, &resp)
	fmt.Println("Cheng", resp)
	err = rpcClient.Call("Calc.Chu", params, &resp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Chu", resp)

	err = rpcClient.Call("Calc.All", params, &resp)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("jia", jia)
	fmt.Println("jian", jian)
	fmt.Println("Resp", resp)
}
