package main

import "fmt"

func main()  {

	chanInt := make(chan int,100)

	for i:=0;i<100;i++ {
		chanInt<- i*2
	}

	// 不关闭管道会报deadlock 错误
	//close(chanInt)

	for v := range chanInt {
		fmt.Println(v)
	}

}
