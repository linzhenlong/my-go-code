package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		num, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(num)
	}
}
