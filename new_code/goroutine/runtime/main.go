package main

import (
	"fmt"
	"runtime"
)

func main()  {

	// 获取当前系统CPU个数
	num := runtime.NumCPU()
	fmt.Printf("当前系统CPU个数%d\n", num)

	// 设置num-1的cpu运行go程序
	// go1.8后,默认让程序运行在多个核上,可以不用设置
	// go1.8前，还是要设置一下,可以更高效的利用cpu
	runtime.GOMAXPROCS(num)

}
