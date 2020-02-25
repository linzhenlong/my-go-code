package main

import (
	"fmt"
)

func calc(ch, res chan int, exit chan bool) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			res <- v
		}
	}
	exit <- true
}

func main() {
	chanInt := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)
	go func() {
		for i := 0; i < 1000; i++ {
			chanInt <- i
		}
		close(chanInt)
	}()
	for i := 0; i < 8; i++ {
		go calc(chanInt, resultChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resultChan)
	}()
	//go func() {
	for {
		if v, ok := <-resultChan; !ok {
			break
		} else {
			fmt.Println(v)
		}
	}
	//}()

	//time.Sleep(time.Second * 10)
}
