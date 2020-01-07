package main

import "fmt"

import "sync"

var wg sync.WaitGroup
func hello(i int) {
	fmt.Println("hello",i)
	wg.Done() // 通知计数器减一
}

func main() {
	
	// wg.Add(10000) 
	for i:=0;i<10000;i++ {
		wg.Add(1) // 计数牌加1
		//go hello(i) // 开启独立goroutine去执行hello()

		// 匿名函数方式
		go func(i int) {
			fmt.Println("hello gorutine",i)
			wg.Done()
		}(i)
	}
	fmt.Println("hello main")
	wg.Wait() // 等待done信号之后才推出.
}