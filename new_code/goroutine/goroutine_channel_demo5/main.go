package main

import (
	"fmt"
	"time"
)

// 计算1-到20000之间的所有素数

func InitData(initData chan int, initExitChan chan bool, num int) {
	for i:=(num-1)*10000;i<=num*10000;i++ {
		initData<- i
	}
	initExitChan<- true
}

// 取素数
func PrimeData(initData chan int, primeData chan int, exitChan chan bool) {
	for {
		num ,ok := <-initData

		if !ok {
			break
		}
		flag := true // 是否是素数
		if num == 1 {
			flag = false
		}
		for i:=2;i<num;i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeData<- num
		}
	}
	exitChan<- true
}



func main()  {

	initData := make(chan int, 80000) // 有读有写初始化50够用了
	primeData := make(chan int , 80000) // 先给1000 吧，我也不知道有多少个素数
	
	exitChan := make(chan bool, 8)

	initExitChan := make(chan bool, 8)

	start := time.Now().Unix()

	for i:=1;i<=8;i++ {
		go InitData(initData, initExitChan, i)
	}


	for i:=1;i<=8;i++ {
		go PrimeData(initData, primeData, exitChan)
	}
	go func() {
		for i:=1;i<=8;i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("执行时间",end - start)
		close(primeData)
	}()
	go func() {
		for i:=1;i<=8;i++ {
			<-initExitChan
		}
		close(initData)
	}()

	//var primeSlice []int

	for {
		_, ok := <- primeData
		if !ok {
			break
		}
		//primeSlice = append(primeSlice, num)
	}
	//fmt.Println(primeSlice)
}