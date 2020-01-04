package main

import "sync"

import "fmt"

import "time"

var wg sync.WaitGroup

func Worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second * 1)
	}

	// 如何接收外部命令退出
	wg.Done()
}

func main() {
	wg.Add(1)
	go Worker()
	// 如果优雅的结束子goroutine
	wg.Wait()
	fmt.Println("over")

}