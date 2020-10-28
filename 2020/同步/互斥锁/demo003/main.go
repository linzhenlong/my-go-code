package main

import "sync"

import "log"

import "sync/atomic"

var wg sync.WaitGroup

func add(x *int32) {
	for i := 0; i < 5000; i++ {
		atomic.AddInt32(x, 5)
	}
	wg.Done()
}

func main() {
	var x int32
	wg.Add(5)
	go add(&x)
	go add(&x)
	go add(&x)
	go add(&x)
	go add(&x)
	wg.Wait()
	log.Println("x=", x)
}
