package main

import (
	"fmt"
)

func sendData(ch chan string, exitChan chan struct{}) {
	ch <- "张三"
	ch <- "李四"
	ch <- "王五"
	ch <- "888"
	ch <- "999"
	ch <- "1000"
	ch <- "1000"
	close(ch)
	exitChan <- struct{}{}
}

func main() {
	ch := make(chan string, 5)

	exitChan := make(chan struct{}, 2)
	go sendData(ch, exitChan)

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(v)
		}
		exitChan <- struct{}{}
	}()

	for i := 0; i < 2; i++ {
		<-exitChan
	}
}
