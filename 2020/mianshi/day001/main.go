package main

import "fmt"

func main() {
	deferCall()
}

// 下面这段代码输出的内容
func deferCall() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	
	defer func() {
		fmt.Println("打印后")
	}()
	
	panic("出错了")
}

// 输出:
//	打印后
//	打印中
//	打印前
//	panic: 出错了

// 参考解析：defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic
