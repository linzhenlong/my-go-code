package main

import "sync"

import "fmt"

var wg sync.WaitGroup


func st(i int, c chan int) {
	c<- i
	wg.Done()
}
func te(c, g chan int) {
	for {
		v,ok := <-c
		if !ok {
			break
		}
		g <-v
	}
	close(g)
}

func main() {
	c := make(chan int , 30)
	g := make(chan int, 30)


	for i:=0;i<5;i++ {
		wg.Add(1)
		go st(i, c)
	}
	wg.Wait()
	close(c)
	go te(c, g)

	for {
		if v,ok := <-g;!ok {
			return
		} else {
			fmt.Println(v)
		}
	}
}