package main

import (
	"fmt"
	"time"
)

const TimeFormat  = "2006/01/02 15:04:05"

func main()  {

	// 获取当前时间的纳秒数
	start := time.Now().UnixNano()
	now := time.Now()
	// 获取当前时间并格式化
	fmt.Println(now.Format(TimeFormat))
	test()
	end := time.Now().UnixNano()
	run := end - start
	// 纳秒除1000=微妙，微妙除1000=毫秒
	fmt.Println("run time:",run/(1000*1000))
}

func test()  {
	// sleep 100ms
	fmt.Println("休息一下")
	time.Sleep(time.Millisecond * 200)
}