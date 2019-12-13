package main

import "fmt"

// 起多协程计算1到50每个数的累加数，然后输出

func InitData(initDataChan chan int,n int)  {
	for i:=1;i<=50;i++ {
		initDataChan<- i
	}

	// 初始化数据写完之后关闭管道
	close(initDataChan)
}

func GetRes(initData chan int, resChan chan map[int]int, exitChan chan bool)  {
	for {
		res := 0
		v , ok := <-initData
		if !ok {
			break
		}
		for i:=0;i<=v;i++ {
			res += i
		}
		resMap := make(map[int]int, 1)
		resMap[v] = res
		resChan<- resMap
	}
	exitChan<- true
}

func main()  {

	n := 50
	var initDataChan chan int
	initDataChan = make(chan int, n)

	resChan := make(chan map[int]int, n)

	// 4 个协程 计算结果
	exitChan := make(chan bool,4)

	go InitData(initDataChan, n)

	// 起四个协程
	for i:=1;i<=4;i++ {
		go GetRes(initDataChan,resChan,exitChan)
	}


	// 协程去读exit
	go func() {
		for i:=1;i<=4;i++{
			<-exitChan
		}
		close(resChan)
	}()

	for {
		res ,ok := <-resChan
		if !ok {
			break
		}
		fmt.Printf("%v\n", res)
	}

}
