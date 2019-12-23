package main

import "fmt"

func test(c chan int)  {
	for i:=0;i<100;i++ {
		if i % 10 == 0 {
			fmt.Println(i)
		}
	}
	c <- 1
}
func test1()  {
	for i:=0;i<100;i++ {
		if i % 10 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	//message := make(chan string)
	c := make(chan int)
	/*go func() {
		message <- "hello goroutine"
		//fmt.Println("hello goroutine")
	}()*/

	//go test1()
	go test(c)
	//fmt.Println(<- message)
	fmt.Println("hello world")
	fmt.Println(<-c)
}
