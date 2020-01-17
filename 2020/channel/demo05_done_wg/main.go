package main

import "fmt"

import "sync"


type Worker struct {
	in chan int
	done func()
}

func DoWork(id int, w Worker) {
	for  {
		n , ok := <-w.in
		if !ok {
			break
		}
		fmt.Printf("Worker %d,received %c \n", id, n)
		// done <- true
		w.done()
	}

	// for range 也是可以滴
	/* for n := range c {
		fmt.Printf("Worker %d,received %d \n", id, n)
	}  */
}

func CreateWorker(id int ,wg *sync.WaitGroup) Worker {
	w := Worker{
		in : make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go DoWork(id, w)
	return w
}

func chanDemo() {
	var wokers [10]Worker
	var wg sync.WaitGroup
	
	for i:=0;i<10;i++ {
		wokers[i] = CreateWorker(i,&wg)
	}
	wg.Add(20)
	for i, worker := range wokers {
		worker.in <- 'a' + i
	}
	for i, worker := range wokers {
		worker.in <- 'A' + i
	}
	wg.Wait()

	

	// 等待全部结束，在退出

	
}

func main() {
	chanDemo()
}
//该看 10-3 使用Channel进行树的遍历