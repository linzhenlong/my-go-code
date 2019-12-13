package main

import "fmt"

//定义一个cat结构体,将cat的各个字段/属性信息，
//放入到cat 结构体中

type Cat struct {
	Name string
	Age  int
	Color string
	Hobby string
}

func main()  {

	// 创建一个cat 变量
	var cat1 Cat

	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "绿色"
	cat1.Hobby = "钓鱼"
	fmt.Println(cat1.Hobby)
}