package main

import "fmt"

// 管道关闭
func main()  {
	var intChan chan int
	intChan = make(chan int ,3)
	intChan <- 1
	intChan <- 2
	close(intChan)
	// 关闭的管道就不能在继续往里面写了，否则会报错

	// panic: send on closed channel
	//intChan <- 3

	fmt.Println("ok")
	// 当管道关闭后，可以正常读取
	fmt.Println(<-intChan)
}
