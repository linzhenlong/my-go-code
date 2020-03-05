package main

import(
	"fmt"
	_"time"

)

type worker struct {
	in chan int
	done chan bool
}
func doWorker(id int, w worker) {
	for {
		n, ok := <-w.in
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %c\n",id, n)
		/* go func() {
			w.done <- true
		}() */
		w.done <- true
	}
}

func createWorker(id int) worker{
	w := worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w)
	return w
}

func channelDemo() {
	var workers [10]worker
	for i:=0;i<10;i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i 
	}
	for _, worker := range workers {
		<- worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	for _, worker := range workers {
		<- worker.done
	}
	//time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
}