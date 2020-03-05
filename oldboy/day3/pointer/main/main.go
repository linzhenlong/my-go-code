package main

import "fmt"

func main()  {
	var i int = 100

	fmt.Println("变量i的地址",&i) //i的地址 0xc420084008
	// 声明一个指针
	var p *int
	p = &i
	fmt.Println(*p)
	*p = 1000
	fmt.Println(i)
	test(&i)

}

func test(p *int)  {
	fmt.Println(p)
	*p = 100
	return
}