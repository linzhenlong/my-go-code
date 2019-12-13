package main

import (
	"fmt"
	"strconv"
)

func main() {

	//1.定义一个容量为10 的管道 int 类型的
	var chan1 chan int
	chan1 = make(chan int, 10)

	for i := 1; i <= 10; i++ {
		chan1 <- i
	}
	// 2.顶一个容量为50 的管道 string 类型的
	var chan2 chan string
	chan2 = make(chan string, 5)

	for i := 1; i <= 5; i++ {
		chan2 <- "fuck" + strconv.Itoa(i)
	}
	// 传统方式在遍历管道时，如果不关闭管道会导致死锁

	// 问题:在实际开发中,可能不知道什么时候才可以关闭管道

	// 可以使用select 方式解决
	start:
	for {
		select {
		// 注意:如果管道一直没有关闭，不会一直阻塞导致死锁，会自动到下一个case去取
		case v := <-chan1:
			fmt.Printf("从chan1读取数据:%d\n", v)
		case v := <-chan2:
			fmt.Printf("从chan1读取数据:%s\n", v)
		default:
			fmt.Println("都取不到了")
			break start // break 到start处
		}
	}

}
