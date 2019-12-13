package main

import (
	"fmt"
	"strconv"
)

//计算1到50的阶乘，并将结果放到map中输出出来

func GetResult(chanNum chan int,chanResult chan map[string]uint64)  {
	for num := range chanNum {
		var res uint64
		res = 1
		var resMap map[string]uint64
		resMap = make(map[string]uint64, 1)
		for i:=1;i<=num;i++ {
			res *= uint64(i)
		}
		resMap[strconv.Itoa(num)+"!"] = res
		chanResult<- resMap
	}
	close(chanResult)
}
func main()  {

	var chanNum chan int
	chanNum = make(chan int, 50)

	chanResult := make(chan map[string]uint64, 50)

	existChan := make(chan bool, 1)


	// init 数据
	go func() {
		for i:=1;i<=50 ;i++  {
			chanNum<- i
		}
		close(chanNum)
	}()

	// 生成结果
	go GetResult(chanNum,chanResult)

	go func() {
		for {
			v, ok := <-chanResult
			if !ok {
				break;
			}
			fmt.Printf("%v\n",v)
		}
		existChan<- true
		close(existChan)
	}()

	// 阻塞一下主进程
	for {
		if <-existChan {
			break
		}
	}
}
