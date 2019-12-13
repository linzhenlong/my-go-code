package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func goroutine()  {
	for i:=1;i<=10;i++{
		fmt.Println("goroutine() hello,world"+strconv.Itoa(i)+"--"+strconv.Itoa(os.Getppid()))
		time.Sleep(time.Second)
	}
}

func main()  {
	// 开启一个协程
	go goroutine()
	for i:=1;i<=10;i++{
		fmt.Println("main() hello ,golang"+strconv.Itoa(i)+"--"+strconv.Itoa(os.Getpid()))
		time.Sleep(time.Second)
	}
}
