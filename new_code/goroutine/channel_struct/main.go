package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func main()  {
	var chanStruct chan Cat
	chanStruct = make(chan Cat, 3)

	cat1 := Cat{Name:"小花猫",Age:10}
	chanStruct <- cat1

	cat2 := Cat{Name:"白花猫",Age:11}
	chanStruct <- cat2

	fmt.Println(<-chanStruct)
	fmt.Println(<-chanStruct)

	// 指针类型
	chanStructZhiZhen := make(chan *Cat, 3)
	cat3 := Cat{Name:"白花猫指针",Age:11}
	chanStructZhiZhen <- &cat3
	fmt.Println(<-chanStructZhiZhen)

}
