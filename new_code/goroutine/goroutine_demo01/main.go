package main

import (
	"fmt"
)

func test(i int, c chan int) {
	fmt.Println(i)
	c <- i
}

func main() {
	c := make(chan int, 10)
	for i := 0; i <= 100; i++ {
		go test(i, c)
	}

	for i := 0; i <= 100; i++ {
		<-c
	}

}
