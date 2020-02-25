package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	ch2 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch2 <- i
	}

	for {
		select {
		case v := <-ch:
			fmt.Println("ch1:", v)
		case v := <-ch2:
			fmt.Println("ch2:", v)
		default:
			goto LABLE
		}
	}
LABLE:
}
