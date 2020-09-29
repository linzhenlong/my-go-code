package main

import "fmt"

import "time"

func main() {
	//  有缓冲通道
	ch := make(chan int, 5)
	fmt.Printf("len(ch) = %d,cap(ch)=%d\n", len(ch), cap(ch))
	go func() {
		defer fmt.Println("子协程退出了")
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("子协程[%d]:len(ch) = %d,cap(ch)=%d\n", i, len(ch), cap(ch))
		}
	}()

	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Printf("主协程[%d]:len(ch) = %d,cap(ch)=%d\n", num, len(ch), cap(ch))
	}
	fmt.Println("main over")
	time.Sleep(time.Second * 2)
}
