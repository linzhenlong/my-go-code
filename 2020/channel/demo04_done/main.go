package main

import "fmt"


type Worker struct {
	in chan int
	done chan bool
}

func DoWork(id int, c chan int, done chan bool) {
	for  {
		n , ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d,received %c \n", id, n)
		// done <- true
		go func() {
			done <- true
		}()
	}

	// for range 也是可以滴
	/* for n := range c {
		fmt.Printf("Worker %d,received %d \n", id, n)
	}  */
}

func CreateWorker(id int ) Worker {
	w := Worker{
		in : make(chan int),
		done: make(chan bool),
	}
	go DoWork(id, w.in,w.done)
	return w
}

func chanDemo() {
	var wokers [10]Worker
	for i:=0;i<10;i++ {
		wokers[i] = CreateWorker(i)
	}

	for i, worker := range wokers {
		worker.in <- 'a' + i
	}
	for i, worker := range wokers {
		worker.in <- 'A' + i
	}

	for _, worker := range wokers {
		// 因为发了两遍worker 因此要收两遍
		<- worker.done
		<- worker.done
	}

	// 等待全部结束，在退出

	
}

func main() {
	chanDemo()
}