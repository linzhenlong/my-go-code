package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求:现在要计算1-200的各个数的阶乘，
// 并且把各个数的阶乘放到map中,最后显示出来

// 思路
// 1.编写一个函数计算各个数的阶乘，并放到map中
// 2. 我们启动的协程是多个，我们统一将多个阶乘的结果放入到map中
// 3.map 应该做出一个全局的

var resultMap = make(map[int]uint64, 10)

// 声明一个全局的互斥锁

// lock 是一个全局变量，全局互斥锁
// sync 是个包
// Mutex 是互斥锁
var lock sync.Mutex

func GoRoutine(n int)  {
	var res uint64
	res = 1
	for i:=1;i<=n;i++ {
		res *= uint64(i)
	}
	// 写map之前加锁
	lock.Lock()
	resultMap[n] = res // 问题1:fatal error: concurrent map writes 同时写一个map
	// 写完解锁
	lock.Unlock()

}

func main()  {
	for i:=1;i<=20;i++  {
		go GoRoutine(i)
	}
	// 问题2:可能会出现主线程先结束，协程不会继续执行
	time.Sleep(5 * time.Second)
	lock.Lock()
	for index, v := range resultMap {
		fmt.Printf("map[%d]=%d\n",index, v)
	}
	lock.Unlock()
}
