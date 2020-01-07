package main

import "fmt"

import "runtime"

import "sync"

var wg sync.WaitGroup

func a() {
	for i:=0;i<10;i++ {
		fmt.Println("A",i)
	}
	wg.Done()
}
func b() {
	for i:=0;i<10;i++ {
		fmt.Println("b",i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(3) // 只是使用一个cpu
	wg.Add(2)
	go a()
	go b()
	wg.Wait()

	
}