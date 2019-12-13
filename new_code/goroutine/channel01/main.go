package main

import "fmt"

// 需求:现在要计算1-200的各个数的阶乘，
// 并且把各个数的阶乘放到map中,最后显示出来

// 思路
// 1.编写一个函数计算各个数的阶乘，并放到map中
// 2. 我们启动的协程是多个，我们统一将多个阶乘的结果放入到map中
// 3.map 应该做出一个全局的

var resultMap = make(map[int]int, 10)

func GoRoutine(n int)  {
	res := 1
	for i:=1;i<=n;i++ {
		res *= i
	}
	resultMap[n] = res // 问题1:fatal error: concurrent map writes 同时写一个map
}

func main()  {
	for i:=1;i<=100;i++  {
		go GoRoutine(i)
	}
	// 问题2:可能会出现主线程先结束，协程不会继续执行
	for index, v := range resultMap {
		fmt.Printf("map[%d]=%d\n",index, v)
	}
}
