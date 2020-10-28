package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify bool

// 为什么需要context
func f() {
	defer wg.Done()
	for {
		fmt.Println("this is f()")
		time.Sleep(time.Millisecond * 100)
		if notify {
			break
		}
	}

}

func main() {
	wg.Add(1)
	go f()
	// 如果通知子goroutine退出
	time.Sleep(time.Second * 2)
	notify = true
	wg.Wait()
}
