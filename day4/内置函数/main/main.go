package main

import (
	"errors"
	"fmt"
	rand2 "math/rand"
	"time"
)

func main()  {
	var i int
	j := new(int)
	*j = 100 //指针赋值
	fmt.Println(i) // 0
	fmt.Println(j) // 内存地址
	fmt.Println(*j) // 0

	var arr []int
	rand2.Seed(time.Now().UnixNano())
	for i:=0;i<=5;i++ {
		randNum := rand2.Intn(10000)
		arr = append(arr,randNum)
	}
	fmt.Println(arr)
	arr = append(arr,arr...) // 追加一个切片
	fmt.Println(arr)

	testPanic()

	newOrMake()
}

func testPanic()  {

	// 捕获异常
	defer func() {
		if err := recover();err != nil {
			fmt.Println(err)
		}
	}()

	initerr := initConfig()
	if initerr != nil {
		panic(initerr)
	}
	b := 0
	a := 100 / b
	fmt.Println(a)
	return
}

func initConfig() (err error)  {
	return errors.New("init config error")
}

func newOrMake() {
	s1 := new([]int)

	fmt.Println(s1)
	*s1 = make([]int,5)
	(*s1)[0] = 99
	fmt.Println(s1)

	s2 := make([]int,5)

	s2[0] = 1000
	fmt.Println(s2)
}