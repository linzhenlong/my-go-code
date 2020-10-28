package main

import "log"

import "time"

func main() {
	// select 判断管道是否存满东西

	chanInt := make(chan int, 10)

	// 子协程写数据
	go func() {
		for {
			select {
			case chanInt <- 1:
				log.Println("write")
			default:
				log.Println("管道写满了")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 取数据
	for val := range chanInt {
		log.Println(val)
		time.Sleep(time.Second * 1)
	}
}
