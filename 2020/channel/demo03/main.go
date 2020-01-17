package main

import "fmt"

import "time"

func Worker(id int, c chan int) {
	/* for  {
		n , ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d,received %d \n", id, n)
	} */

	// for range 也是可以滴
	for n := range c {
		fmt.Printf("Worker %d,received %d \n", id, n)
	} 
}

func CreateWorker(id int ) chan<- int {
	c := make(chan int)
	go Worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i:=0;i<10;i++ {
		channels[i] = CreateWorker(i)
	}
	for i:=0;i<10;i++ {
		channels[i] <- 'a' + i
	}
	for i:=0;i<10;i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Second)
	
}


func bufferedChannel() {
	c := make(chan int, 3)
	go Worker(0, c)
	c <- 'a'+ 1
	c <- 'a'+ 2
	c <- 'a' + 3
	c <- 'a' + 4
	time.Sleep(time.Second)
}

func closeChannel() {
	// 有明确结尾的可以加close
	// 一般从close 都是由发送方close的
	c := make(chan int)
	go Worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Second)
	

}
func main() {
	//chanDemo()
	//bufferedChannel()
	closeChannel()
}