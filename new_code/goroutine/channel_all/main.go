package main

import "fmt"

type Cat struct {
	Name string
	Age int
}


func main()  {

	// 管道放入任意类型
	var channelAll chan interface{}
	channelAll = make(chan interface{}, 1)

	// 放入字符串
	channelAll <- "hello world"
	fmt.Printf("放入字符串，在取出:%v\n", <- channelAll)
	// 放入int
	channelAll <- 10
	fmt.Printf("放入int，在取出:%v\n", <- channelAll)

	// 放入数组
	channelAll <- [3]int{1,3,4}
	fmt.Printf("放入数组，在取出:%v\n", <- channelAll)

	// 放入切片
	channelAll <- []string{"张三","李四","王五"}
	fmt.Printf("放入切片，在取出:%v\n", <- channelAll)

	// 放入map
	myMap := make(map[string]string,10)
	myMap["name1"] = "张三"
	channelAll <- myMap
	fmt.Printf("放入map，在取出:%v\n", <- channelAll)
	myMap["name2"] = "李四"
	channelAll<- myMap
	fmt.Printf("放入map，在取出:%v\n", <- channelAll)
	// 结构体
	cat := Cat{
		Name:"小花猫",
		Age:18,
	}
	channelAll <- cat
	cat1 := <- channelAll

	fmt.Printf("cat1 的类型=%T,值:%v\n",cat1,cat1)

	// cat1.Name undefined (type interface {} is interface with no methods)
	// 会报错需要使用类型断言
	// fmt.Printf("cat1.Name=%v\n",cat1.Name)
	 fmt.Printf("cat1.Name=%v\n",cat1.(Cat).Name)


	// 放入结构体指针
	channelAll <- &cat
	fmt.Printf("放入结构体指针，在取出:%v\n", <- channelAll)
}
