package main

import "sync"

import "fmt"

import "time"

var wg sync.WaitGroup

// 全局变量控制是否退出
var exit bool

// 全局变量方式存在的问题
// 1.使用全局变量在跨包调用时不容易统一
// 2.如果Process中再启动goroutine，就不太好控制了.

func Process() {
	for {
		fmt.Println("process run...")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go Process()
	// sleep 3秒以免程序过快退出
	time.Sleep(time.Second * 3)
	// 修改全局变量实现子goroutine的退出
	exit = true
	wg.Wait()
	fmt.Println("main over")

}
