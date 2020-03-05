package main

import(
	"fmt"
	"sync"
	_"time"

)

type worker struct {
	in chan int
	done func()
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
		w.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker{
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func channelDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i:=0;i<10;i++ {
		workers[i] = createWorker(i, &wg)
	}
	
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i 
	}
	
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
	//time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
}