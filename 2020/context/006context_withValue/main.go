package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// TraceCode .
type TraceCode string

func worker(ctx context.Context) {
	key := TraceCode("trace_code")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("worker traceCode invalid")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(10 * time.Millisecond)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
	wg.Done()
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	ctx = context.WithValue(ctx, TraceCode("trace_code"), "123-321-1234567")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()

}
