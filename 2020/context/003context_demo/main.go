package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("this is f()")
		log.Print("this is f()")
		time.Sleep(time.Millisecond * 600)
		select {
		case <-ctx.Done():
			log.Print("LOOP")
			break LOOP
		default:
			log.Print("default")
		}
	}
}

func main() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go f(ctx)
	/* time.Sleep(time.Second * 3)
	cancel() */
	wg.Wait()
}
