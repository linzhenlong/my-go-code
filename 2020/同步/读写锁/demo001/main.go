package main

import "sync"

import "log"

import "time"

import "fmt"

var (
	// 声明读写锁.
	rwLock sync.RWMutex
	wg     sync.WaitGroup
	// 全局变量
	x int
)

func write() {
	//锁一下
	rwLock.Lock()
	log.Println("write rwlock")
	x++
	time.Sleep(2 * time.Second)
	log.Println("write rwunlock")
	// 解锁
	rwLock.Unlock()
	wg.Done()
}

func read(i int) {

	// 读锁
	rwLock.RLock()
	fmt.Println("read rwlock")
	log.Printf("goroutine:%d,x=%d\n", i, x)
	time.Sleep(2 * time.Second)
	log.Println("read rwunlock")
	// 解锁
	rwLock.RUnlock() // 读解锁

	wg.Done()
}

func main() {
	wg.Add(1)
	go write()
	time.Sleep(time.Millisecond * 5)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read(i)
	}
	wg.Wait()
}
