package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	for {
		if val, ok := <-ch; ok {
			fmt.Println(val)
		} else {
			break
		}
	}
}
