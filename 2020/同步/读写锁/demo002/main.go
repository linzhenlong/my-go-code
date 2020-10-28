package main

import "sync"

import "time"

import "fmt"

// 互斥锁与读写锁性能对比
var (
	// 互斥锁
	mutexLock sync.Mutex
	rwLock    sync.RWMutex
	wg        sync.WaitGroup
	// 全局变量
	x int
)

// 写数据
func write() {
	for i := 0; i < 100; i++ {
		//mutexLock.Lock()
		rwLock.Lock()
		x++
		time.Sleep(1 * time.Millisecond)
		//mutexLock.Unlock()
		rwLock.Unlock()
	}
	wg.Done()
}

func read(i int) {
	for i := 0; i < 100; i++ {
		//mutexLock.Lock()
		rwLock.RLock()
		time.Sleep(1 * time.Millisecond)
		//mutexLock.Unlock()
		rwLock.RUnlock()
	}
	wg.Done()
}

// 互斥锁: 12955336000纳秒
// 读写锁: 271579000
func main() {
	start := time.Now().UnixNano()
	wg.Add(1)
	go write()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read(i)
	}
	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Println("run time:", end-start)
}
