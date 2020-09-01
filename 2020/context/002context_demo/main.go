package main

import (
	"fmt"
	"time"
)

var ch = make(chan struct{})

func f() {
LOOP:
	for {
		fmt.Println("this is f()")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ch:
			fmt.Println("loop")
			break LOOP
		default:
		}
	}
}

func main() {
	go f()
	time.Sleep(time.Second * 1)
	ch <- struct{}{}
}
