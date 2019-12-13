package main

import (
	"fmt"
	"time"
)

func WriteData(chanInt chan int)  {
	for i:=0;i<500;i++ {
		chanInt<- i
		fmt.Printf("WriteData chanInt<- %d\n", i)
	}
	// 写完关闭管道
	close(chanInt)
}

func ReadData(chanInt chan int, isFinished chan bool)  {
	for v := range chanInt{
		fmt.Printf("ReadData <-chanInt %d\n", v)
		time.Sleep(time.Millisecond * 10)
	}
	isFinished<- true
	close(isFinished)
}

func main()  {

	var chanInt chan int

	// 用于阻塞主线程
	isFinished := make(chan bool, 1)

	chanInt = make(chan int, 5)
	go WriteData(chanInt)

	go ReadData(chanInt,isFinished)

	// 阻塞主进程，知道finish 完成了,在退出
	for {
		_, ok := <-isFinished
		if ok {
			break
		}
	}

}
