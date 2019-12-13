package main

import "fmt"

func main()  {

	var t float32
	var x interface{}
	x = t
	// 转成float32
	y := x.(float32)
	// 转成float32 代检测的
	z, ok := x.(float32)
	if ok {
		fmt.Println("success",z)
	} else {
		fmt.Println("error")
	}
	fmt.Println(y)
}


