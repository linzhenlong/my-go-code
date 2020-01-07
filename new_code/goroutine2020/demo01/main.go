package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func main() {
	go hello() // 开启独立goroutine去执行hello()
	fmt.Println("hello main")
}