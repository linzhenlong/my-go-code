package main

import "runtime"

import "fmt"

func main() {

	runtime.GOMAXPROCS(3)
	for { 
		go fmt.Println(0)
		fmt.Println(1)
	}

}
