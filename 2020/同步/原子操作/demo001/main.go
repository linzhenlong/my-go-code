package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	x     int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

// 互斥锁操作
func add1() {
	for i := 0; i < 500; i++ {
		mutex.Lock()
		x++
		mutex.Unlock()
	}
	wg.Done()
}

// 互斥锁操作：57811000纳秒
func add2() {
	for i := 0; i < 500; i++ {
		atomic.AddInt32(&x, 1)
	}
	wg.Done()
}

// 原子操作:10315000纳秒
func main() {
	start := time.Now().UnixNano()
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add1()
	}
	
	wg.Wait()
	fmt.Println(x)
	end := time.Now().UnixNano()
	fmt.Println("run time:", end-start)
}
