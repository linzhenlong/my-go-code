package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(1)
	for {
		intChan := make(chan int, 1)
		stringChan := make(chan string, 1)
		intChan <- 1
		//stringChan <- "string"
		select {
		case value := <-intChan:
			fmt.Println(strconv.Itoa(value) + " out put")
		case value := <-stringChan:
			panic(value + " out put")
		}
	}
}
