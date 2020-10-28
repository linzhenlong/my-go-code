package main

import (
	"fmt"
	"sync"
)

// map不是并发安全的，但是并发读没有问题
// map的并发写是有问题的
// 1.加锁解决
// 2.sync.Map

func main() {
	wg := sync.WaitGroup{}
	m := make(map[int]int)
	mutex := sync.Mutex{}

	// 并发写入数据
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			// 1.加锁解决
			mutex.Lock()
			m[i] = i // fatal error: concurrent map writes
			mutex.Unlock()
		}(i)
	}
	
	wg.Wait()
	fmt.Println("m:",m)
	fmt.Println("main over")
}
