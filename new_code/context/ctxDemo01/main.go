package main

import "sync"

import "context"

import "time"

import "fmt"


var wg sync.WaitGroup

func Process(ctx context.Context) {
	LOOP:
	for {
		fmt.Println("process run...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():// 等待上级通知
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
	time.Sleep(time.Second * 4)

	// 通知子goroutine结束
	cancel() 
	wg.Wait()
	fmt.Println("main over")
} 