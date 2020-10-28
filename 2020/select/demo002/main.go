package main

import "log"

func main() {
	// 2个管道
	outChan1 := make(chan int, 1)
	outChan2 := make(chan string, 1)
	outChan3 := make(chan string, 1)

	go func() {
		//time.Sleep(time.Second * 2)
		outChan1 <- 666
	}()
	go func() {
		//time.Sleep(time.Second * 4)
		outChan2 <- "777"
	}()
	go func() {
		//time.Sleep(time.Second * 4)
		outChan3 <- "888"
	}()

	// 多个channel同时ready 随机执行一个
	select {
	case val1 := <-outChan1:
		log.Println(val1)
	case val2 := <-outChan2:
		log.Println(val2)
	case val3 := <-outChan3:
		log.Println(val3)
	}
	log.Println("main over")
}
