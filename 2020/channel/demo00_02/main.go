package main

import "fmt"

import "time"

func main() {
	// 0代表无缓冲通道.
	ch := make(chan int, 0)
	fmt.Printf("len(ch)=%d,cap(len)=%d\n", len(ch), cap(ch))

	// 子协程写数据
	go func() {
		defer fmt.Println("子协程结束了")
		for i := 0; i <= 5; i++ {
			ch <- i
			fmt.Printf("子协程在运行[%d]:len(ch)=%d,cap(len)=%d\n", i, len(ch), cap(ch))
		}
	}()
	time.Sleep(time.Second * 2)
	for i := 0; i <= 5; i++ {
		num := <-ch
		fmt.Printf("主协程在运行[%d]:len(ch)=%d,cap(len)=%d\n", num, len(ch), cap(ch))
	}
	fmt.Println("main 结束")
}
