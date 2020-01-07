package main

import "fmt"

var name string = "张三"
func test() {
	name := "tom"
	age :=19
	fmt.Println("func test", name,age)
}

func main() {
	test()
	fmt.Println(name) // 输出为张三，name虽然定义了全局变量，并且在test函数中赋值了，但是，test 中赋值作用域为test函数

	i := 18
	for i=0;i<5;i++ {
		fmt.Println("i",i)
	}
	fmt.Println("i:=18的i==>",i)
}