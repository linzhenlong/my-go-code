package main

import "log"

import "sync"

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {

		log.Println("子协程1执行了")
		wg.Done()
	}()
	go func() {

		log.Println("子协程2执行了")
		wg.Done()

	}()
	go func() {
		log.Println("子协程3执行了")
		wg.Done()

	}()
	wg.Wait()
}
