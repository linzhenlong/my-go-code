package main

import "sync"

import "log"

import "time"

func main() {
	ch := make(chan int, 10)
	var i int
	var wg sync.WaitGroup
	for {
		i++
		if i > 100 {
			break
		}

		wg.Add(1)
		log.Println("wg.add")

		go func() {
			ch <- 1
			defer wg.Done()
			time.Sleep(time.Second * 2)
			<-ch
		}()
	}
	wg.Wait()
	log.Println("main over")
}
