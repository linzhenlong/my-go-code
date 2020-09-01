package main

import "time"

var c = make(chan int) // 无缓冲channel

var a int

func f() {
	a = 1
	select {
	case v := <-c:
		print(v)
	default:
		print("defualt")
		time.Sleep(time.Second)
	}

}
func main() {
	go f()
	c <- 0
	print(a)
}
