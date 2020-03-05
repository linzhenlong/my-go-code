package main

import (
	_"fmt"
	"log"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for {
		time.Sleep(time.Second)
		n, ok := <-c
		if !ok {
			break
		}
		//fmt.Printf("Worker %d received %d\n",id, n)
		log.Printf("Worker %d received %d\n",id, n)
	}
}
func createWorker(id int) chan<- int{
	c  := make(chan int)
	go worker(id, c)
	return c
}
func main() {
	var c1, c2 = generator(),generator()
	worker := createWorker(0)
	
	var values []int

	// 10秒之后退出.
	tm := time.After(20 * time.Second)

	// 每秒钟看队列长度
	tick := time.Tick(time.Second) 
	
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
			//fmt.Println("Received from c2:", n)
		case  activeWorker<-activeValue:
			values = values[1:]
		case <-tm:
			log.Printf("Bye")
			return
		case <- time.After(800 * time.Millisecond):
			log.Printf("time out")
		case <- tick:
			log.Printf("len(values):%d", len(values))
		/* default:
			fmt.Println("No value received") */
		}
	}
}
