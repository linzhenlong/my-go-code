package main

import "sync"

import "context"

import "fmt"

import "time"

var wg sync.WaitGroup

func Process2(ctx context.Context) {
	LOOP:
	for {
		fmt.Println("Process2 run...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func Process(ctx context.Context) {
	go Process2(ctx)
	LOOP:
	for {
		fmt.Println("Process run...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go Process(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
	fmt.Println("main over...")

}