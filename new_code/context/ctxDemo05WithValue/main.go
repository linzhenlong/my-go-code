package main

import "sync"

import "context"

import "time"

import "fmt"

type TraceCode string
type UserID string

var wg sync.WaitGroup


func Process2(ctx context.Context) {
	key1 := TraceCode("TRACE_CODE")
	traceCode1, ok := ctx.Value(key1).(string) // 在子goroutine中获取trace code

	key2 := TraceCode("TRACE_CODE2")
	traceCode2, ok := ctx.Value(key2).(string) // 在子goroutine中获取trace code

	if !ok {
		fmt.Println("非法的trace code...")
	}
	LABLE:
	for {
		fmt.Printf("Process2 trace_code1=>%s,trace_code12=>%s\n",traceCode1, traceCode2)
		time.Sleep(time.Millisecond * 5)
		select {
		case <-ctx.Done():
			break LABLE
		default:
		}
	}

}

func Process(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code

	if !ok {
		fmt.Println("非法的trace code...")
	}
	ctx = context.WithValue(ctx,TraceCode("TRACE_CODE2"), "77889966")
	userIDkey := UserID("user_id")
	userID, _ := ctx.Value(userIDkey).(int)
	go Process2(ctx)
	LABLE:
	for {
		fmt.Printf("Process, trace code:%s\n",traceCode)
		fmt.Printf("Process, user_id:%d\n",userID)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LABLE
		default:

		}
	}
	fmt.Println("Process done....")
	wg.Done()
}

func main() {

	// 设置50ms的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 50)

	// 在系统的入口设置trace code 传递给后续启动的goroutine实现日志聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	// 用户ID.
	ctx = context.WithValue(ctx,UserID("user_id"), 9999999)
	wg.Add(1)
	go Process(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine 结束
	wg.Wait()
	fmt.Println("main over...")

}