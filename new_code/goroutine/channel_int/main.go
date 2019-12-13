package main

import "fmt"

func main()  {

	// 声明一个长度为3的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 看channel是什么
	fmt.Printf("intChan的值=%v,地址=%p\n",intChan, &intChan)

	// 3.想管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 4.看看管道的长度(len)及容量(cap)

	fmt.Printf("intChan 的长度：%d,容量:%d",len(intChan),cap(intChan))

	// 注意当我们给管道添加数据是，不能超过他的容量也就是初始化make时的长度，否则会报错
	intChan <- 88


	//intChan <- 77
	// 超过长度报以下错误:
	// fatal error: all goroutines are asleep - deadlock!
	//
	//goroutine 1 [chan send]:

	// 可以取一个放一个
	// 5.从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println(num2)
	intChan <- num2
	num2 = <-intChan
	fmt.Println(num2)

	// 6.在没有使用协程的情况下，我们的管道数据已经全部取出，那么在取也会报错 fatal error: all goroutines are asleep - deadlock!
	num2 = <- intChan
	num2 = <- intChan
	num2 = <- intChan

}