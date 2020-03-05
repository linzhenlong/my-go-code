package main

import (
	"fmt"
	"time"
)

func main()  {
	recusive(0)
	n := calc(5)
	fmt.Println(n)
	for j:=0;j<10;j++  {
		m := fab(j)
		fmt.Println(m)
	}
}

func recusive(n int)  {
	fmt.Println("hello")
	time.Sleep(time.Second)
	if n > 3 {
		return
	}
	recusive(n + 1)
}
// 阶乘
func calc(n int) int  {
	if n==1 {
		return 1
	}
	return calc(n-1) * n
}
// 费布那切数列
func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}