package main

import "sync"

import "log"

var x int
var wg sync.WaitGroup

// 声明一个锁
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	log.Println("x=", x)
}
