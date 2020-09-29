package main

import "time"

import "log"

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func main() {
	// 2个管道
	outChan1 := make(chan string)
	outChan2 := make(chan string)

	// 跑两个子协程，写数据
	go test1(outChan1)
	go test2(outChan2)

	timer := time.NewTimer(time.Second * 6)
	// 用select监控
LOOP:
	for {
		select {
		case s1 := <-outChan1:
			log.Println(s1)

		case s2 := <-outChan2:
			log.Println(s2)
		case <-timer.C:
			log.Println("time over")
			break LOOP
		default:
			time.Sleep(time.Second)
			log.Println("default")
		}

	}
}
