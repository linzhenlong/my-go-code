package main

import "fmt"

// 需求:现在要计算1-200的各个数的阶乘，
// 并且把各个数的阶乘放到map中,最后显示出来

// 思路
// 1.编写一个函数计算各个数的阶乘，并放到map中
// 2. 我们启动的协程是多个，我们统一将多个阶乘的结果放入到map中
// 3.map 应该做出一个全局的

func InitData(size int, initDataChan chan int)  {
	if size < 1{
		return
	}
	for i:=1;i<=size;i++  {
		initDataChan <- i
	}
	close(initDataChan)
}

func GetRes(initDataChan chan int, resChan chan map[int]uint64,exitChan chan bool)  {
	for {
		n, ok := <-initDataChan
		if !ok {
			break
		}
		var res uint64
		res = 1;
		for i := 1; i <= n; i++ {
			res *= uint64(i)
		}
		myMap := make(map[int]uint64, 1)
		myMap[n] = res
		resChan <- myMap
	}
	exitChan <- true
}

func main()  {
	initDataChan := make(chan int, 1000)
	go InitData(200, initDataChan)
	resChan := make(chan map[int]uint64, 200)
	exitChan := make(chan bool,4)

	for i:=0;i<4;i++ {
		go GetRes(initDataChan,resChan,exitChan)
	}
	go func() {
		for i:=0;i<4;i++ {
			<- exitChan
		}
		close(resChan)
	}()

	for  {
		res , ok := <-resChan
		if !ok {
			break
		}
		for index ,v := range res {
			fmt.Printf("%d!=%v\n",index, v)
		}
	}

}
