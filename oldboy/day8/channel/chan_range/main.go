package main

import (
	"fmt"
	"time"
)

func main() {
	chanInt := make(chan int, 100)

	go func() {
		for i := 0; i < 1000; i++ {
			chanInt <- i
		}
		close(chanInt)
	}()

	go func() {
		for v := range chanInt {
			fmt.Println(v)
		}
	}()
	time.Sleep(time.Second * 3)
}
