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
	m := sync.Map{}
	

	// 并发写入数据
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			
			m.Store(i,i)
		}(i)
	}
	wg.Wait()
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key:%d,val:%d\n",key,value)
		return true
	})
	fmt.Println("main over")
}
