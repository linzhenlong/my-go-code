package main

import "log"

func main() {
	ch := make(chan int, 3)

	count := 3
	go func() {

		log.Println("子协程1执行了")
		ch <- 1
	}()
	go func() {

		log.Println("子协程2执行了")
		ch <- 1
	}()
	go func() {
		log.Println("子协程3执行了")
		ch <- 1
	}()
	/* for i := 0; i < 3; i++ {
		log.Println(<-ch)
	} */
	for range ch {
		count--
		if count == 0 {
			close(ch)
		}
	}
}
