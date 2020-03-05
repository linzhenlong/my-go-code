package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("map 没有make 就直接使用")
		}
	}()
	var m map[string]int
	m["stu"] = 100 // panic

}

func sayHello() {
	fmt.Println("hello")
}
func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	for i := 0; i < 10; i++ {
		go test()
		go sayHello()
	}
	time.Sleep(time.Second * 10)
}
