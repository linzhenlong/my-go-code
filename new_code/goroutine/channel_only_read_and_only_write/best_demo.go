package main

import "fmt"

// ch chan<- int 这样ch 就只能写操作了
func send(ch chan<- int, exitChan chan<- interface{})  {
	for i:=0;i<10;i++ {
		ch<- i
	}
	close(ch)

	var a struct{}
	exitChan<- a
}

// ch <-chan int 这样ch 就只能读操作了
func recv(ch <-chan int, exitChan chan<- interface{})  {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan<- a
}

func main()  {

	// channel 只读只写的最佳案例
	var ch chan int
	ch = make(chan int,10)
	exitChan := make(chan interface{},2)
	go send(ch,exitChan)
	go recv(ch, exitChan)

	total := 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("main 结束")

}
