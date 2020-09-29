package main

import "fmt"

import "sync"

// 生产者与消费者

var wg sync.WaitGroup

// 生产者 只写
func producter(data chan<- int) {
	defer close(data) // 关闭管道
	for i := 0; i < 5; i++ {
		data <- i
	}
}

// 消费者只读
func consumer(data <-chan int) {
	for num := range data {
		fmt.Println(num)
	}
	defer wg.Done()
}
func main() {
	ch := make(chan int)
	wg.Add(1)
	// 生产者
	go producter(ch)
	// 消费者
	go consumer(ch)
	wg.Wait()
}
