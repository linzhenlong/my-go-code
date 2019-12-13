package main

import "fmt"

type A struct {
	Num int
}

func main()  {

	var a A
	a.Num = 100
	a.test(90)
	fmt.Println("main a.Num",a.Num)
}

func (a A) test(num int)  {
	a.Num = num
	fmt.Println("test a.Num",a.Num)
}
