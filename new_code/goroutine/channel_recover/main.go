package main

import (
	"fmt"
	"time"
)

func SayFuck()  {
	i :=0
	for {
		i++
		time.Sleep(time.Second)
		fmt.Println("fuck")
		if i > 10 {
			break
		}
	}
}

func Test()  {
	var myMap map[int]string

	// map需要先make 这块故意给他写错了,
	// 会报错 panic: assignment to entry in nil map

	// 这块可以使用错误处理机制defer + recover 来解决,如果没有捕获panic 会影响主线程
	defer func() {
		// 捕获抛出的panic
		if err := recover();err != nil{
			fmt.Println("Test()发生错误了")
		}
	}()
	myMap[0] = "您好" // map需要先make 这块故意给他写错了,

}
func main()  {

	go SayFuck()
	go Test()

	// 阻塞一下主线程
	var a int
	fmt.Scanf("%d",&a)
}
