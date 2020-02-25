package main

import (
	"fmt"
	"log"
	"time"
)

func write(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		log.Println("write data:", i)
	}
}

func read(ch chan int) {
	for {
		if val, ok := <-ch; ok {
			fmt.Println(val)
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	chanInt := make(chan int, 5)
	go write(chanInt)

	go read(chanInt)

	time.Sleep(time.Second * 400)
}
