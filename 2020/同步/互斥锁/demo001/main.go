package main

import "sync"

import "log"

var x int
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x++
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
