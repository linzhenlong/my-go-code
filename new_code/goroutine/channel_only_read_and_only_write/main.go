package main

import "fmt"

func main()  {

	// 管道可以声明未只读或者只写
	// 1.在默认情况下，管道是双向管道，即可读也可写
	var chan1 chan int // 双向的
	chan1 = make(chan int ,10)
	//写入
	chan1<- 10
	//读出
	fmt.Println(<-chan1)

	// 2.声明为只写的
	var chan2 chan<- int
	chan2 = make(chan<- int, 10)
	chan2<- 100

	// 下面代码会报错:invalid operation: <-chan2 (receive from send-only type chan<- int)
	//fmt.Println(<-chan2)

	// 3.声明为只读的
	var chan3 <-chan int

	chan3 = make(<-chan int, 3)

	// 下面代码会报错:invalid operation: chan3 <- 3 (send to receive-only type <-chan int)
	//chan3<- 3

	fmt.Println(chan3)

}
