package main

import "sync"

import "context"

import "time"

import "fmt"

var wg sync.WaitGroup


func Process(ctx context.Context) {
	LOOP:
	for {
		fmt.Println("db connecting....")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done(): //50ms后自动调用
		break LOOP
		default:
		}
	}
	fmt.Println("Process over...")
	wg.Done()
}

func main() {
	// 50ms 超时.
	ctx, cancel := context.WithTimeout(context.Background(),time.Millisecond * 50 ) 
	wg.Add(1)
	go Process(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("main over")
}