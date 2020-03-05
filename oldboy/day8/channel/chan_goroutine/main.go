package main

import (
	"fmt"
	"sync"
)

func sendData(ch chan string) {
	ch <- "张三"
	ch <- "李四"
	ch <- "王五"
	ch <- "888"
	ch <- "999"
	ch <- "1000"
	ch <- "1000"
	close(ch)
}

func main() {
	ch := make(chan string, 5)

	var wg sync.WaitGroup
	wg.Add(1)
	go sendData(ch)

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(v)
		}
		wg.Done()
	}()
	wg.Wait()
}
