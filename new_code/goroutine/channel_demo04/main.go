package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// 创建一个Person结构体{Name,Age}
// 使用rand方法随机创建是个10Person实例，并存放到channel 中
// 变量channel,将各个Person 输出

type Person struct {
	Name string
	Age int
}
func main()  {

	personChan := make(chan Person, 10)

	for i:=1;i<=10 ;i++  {
		personChan <- Person{
			Name:"name"+strconv.Itoa(rand.Intn(10)),
			Age:rand.Intn(10),

		}
	}
	// 需要关闭
	close(personChan)
	for {
		res , ok := <-personChan
		fmt.Printf("%v",ok)
		if !ok {
			return
		}
		fmt.Printf("%v\n", res)
	}
}
