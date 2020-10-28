package main

import (
	"fmt"
	"sync"
)

// map不是并发安全的，但是并发读没有问题
// map的并发写是有问题的

func main() {
	wg := sync.WaitGroup{}
	m := make(map[int]int)

	// 写入数据
	for i := 0; i < 5; i++ {
		m[i] = i
	}
	// 打印数据
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func (i int) {
				fmt.Printf("m[%d]=%d\n", i, m[i])
				wg.Done()
		}(i)
	} 
	wg.Wait()
	fmt.Println("main over")
}
