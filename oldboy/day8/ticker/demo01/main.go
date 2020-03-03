package main

import "time"

import "log"

func queryDb(ch chan int) {
	time.Sleep(time.Millisecond)
	ch <- 1000
}

func main() {
	ch := make(chan int)
	go queryDb(ch)
	t := time.NewTicker(time.Second)
	select {
	case v := <-ch:
		log.Println(v)
	case <-t.C:
		log.Println("timeout")
	}
}
