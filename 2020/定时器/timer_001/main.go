package main

import (
	"fmt"
	"log"
	"time"
)

// Timer:时间到了,执行只执行一次.

func main() {
	// 1.timer 的基本使用
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能响应一次
	/* timer2 := time.NewTimer(time.Second)
	for {
		fmt.Printf("timer2:%v\n", <-timer2.C) // 死锁
	} */

	// 3.timer 实现延时的功能

	// 第一种
	// time.Sleep(time.Second)

	// 第二种
	// timer3 := time.NewTimer(2 * time.Second)
	// log.Println("延时2秒")
	// <-timer3.C
	// log.Println("2秒到了")

	// 第三种
	// log.Println("第三种延时2秒")
	// timeChan := time.After(time.Second * 2)
	// //log.Printf("timeChan:%v\n", <-timeChan)
	// <-timeChan
	// log.Println("2秒到了")

	// 4.停止定时器
	// timer4 := time.NewTimer(time.Second * 3)
	// go func() {
	// 	<-timer4.C
	// 	fmt.Println("定时器执行了")
	// }()
	// b := timer4.Stop()
	// if b {
	// 	fmt.Println("timer4已经关闭")
	// }

	// 5.重置定时器
	timer5 := time.NewTimer(time.Second * 2)
	log.Println("重置前")
	timer5.Reset(time.Second * 3)
	log.Println("重置后")
	<-timer5.C
	log.Println("到时间了")
}
