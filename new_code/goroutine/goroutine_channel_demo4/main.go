package main

import (
	"fmt"
	"sort"
	"strconv"
)


func InitData(initData chan int,n int,finishedInit chan bool)  {
	for i:=(n-1)*100 +1;i<=n*100;i++ {
		initData<- i
	}
	finishedInit<- true
}

func CloseInt(initData chan int, finishedInit chan bool)  {
	for i:=1;i<=10;i++ {
		<-finishedInit
	}
	close(initData)
}

func GetRes(initData chan int, resChan chan map[int]uint64, finishedRes chan bool)  {
	for {
		var res uint64 = 0
		v , ok := <-initData
		if !ok {
			break
		}
		for i:=0;i<=v;i++ {
			res += uint64(i)
		}
		resMap := make(map[int]uint64, 1)
		resMap[v] = res
		resChan<- resMap
	}
	finishedRes<- true
}


func main()  {

	initData := make(chan int , 1000)

	// 10个协程生成数据
	finishedInit := make(chan bool, 10)

	finishedRes := make(chan bool, 100)

	resChan := make(chan map[int]uint64, 1000)

	for i:=1;i<=10;i++ {
		go InitData(initData,i,finishedInit)
	}

	// 通过closeInit关闭initData管道
	go CloseInt(initData, finishedInit)

	// 多协程生成结果
	for j:=1;j<=100;j++ {
		go GetRes(initData, resChan, finishedRes)
	}

	// 关闭resChan管道
	go func() {
		for k:=1;k<=100;k++ {
			<-finishedRes
		}
		close(resChan)
	}()

	var myMap map[int]uint64
	myMap = make(map[int]uint64, 1000)

	var sortSlice []int
	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		for index, value := range v {
			sortSlice = append(sortSlice,index)
			myMap[index] = value
		}
	}
	sort.Ints(sortSlice)
	for _,v := range sortSlice {
		if v > 50 {
			break
		}
		var str string
		for i:=1;i<v;i++ {
			str +=strconv.Itoa(i)+"+"
		}
		fmt.Printf("%s%d=%d\n",str,v,myMap[v])
	}
}
