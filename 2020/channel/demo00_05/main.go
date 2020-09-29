package main

import "fmt"

func main() {
	// 定义一个可读可写的管道
	ch := make(chan int, 3)

	// 将ch 转为只写的管道
	var ch2 chan<- int = ch

	// 将ch转只读的
	var ch3 <-chan int = ch
	ch2 <- 1000        // 写
	fmt.Println(<-ch3) // 读

	// 单向的管道不能在转回去
	//d := (chan int)(ch2)

}
