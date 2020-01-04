package main

import "sync"

import "fmt"

import "time"

var wg sync.WaitGroup

// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

func Process(exitChan chan struct{}) {
	LOOP:
	for {
		fmt.Println("process run...")
		time.Sleep(time.Second)
		select {
		case <-exitChan: // 等待接收上级通知
		break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go Process(exitChan)
	// sleep 3秒以免程序过快退出
	time.Sleep(time.Second*3)

	// 给子goroutine一个退出的信号
	exitChan<- struct{}{} 
	close(exitChan)
	wg.Wait()
	fmt.Println("main over...")
}