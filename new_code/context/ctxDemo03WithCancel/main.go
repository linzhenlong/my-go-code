package main

import "context"

import "fmt"

// 返回只读管道
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for  {
			fmt.Println("n:",n)
			select {
			case <-ctx.Done():
				return // return 结束该goroutine,防止泄露
			case dst<-n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		if n== 5 {
			break // break 结束循环，主进程结束，会走defer cancel()结束上下文
		}
	}
}

