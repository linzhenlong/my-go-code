package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		defer fmt.Println("子协程退出了...")
		fmt.Println("子协程正在运行")
		ch <- 666
	}()
	num := <-ch
	fmt.Println(num)
	fmt.Println("主协程退出")
}
