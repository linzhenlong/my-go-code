package main

import "sync"

import "fmt"


// 下面代码输出什么

// 原因分析:
//原因有二：1，golang是值拷贝传递；
// 2，for循环很快就执行完了，但是创建的10个协程需要做初始化：
// 上下文准备，堆栈，和内核态的线程映射关系的工作，是需要时间的，
// 比for慢，等都准备好了的时候，会同时访问变量temp 。
// 这个时候的i肯定是for执行完成后的数字10。
//所以10个协程都打印10。
 
//破解的方法就是闭包，给匿名函数增加入参，
// 因为是值传递，所以每次for创建一个协程的时候，会拷贝一份i传到这个协程里面去，这样就可以实现0-9的数字打印了
func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i:=0;i<10;i++ {
		
		go func() {
			fmt.Println("i",i) 
			wg.Done()
		}()
	}
	for i:=0;i<10;i++ {
		go func(i int) {
			fmt.Println("j",i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("main over")
}