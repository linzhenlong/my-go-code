package main

import "fmt"

type add_func func(int, int) int

func main()  {
	// 函数赋值给变量
	c := add
	fmt.Println(c)
	sum := c(10,20)
	fmt.Println(sum)

	sum = operator(c,100,200)
	fmt.Println(sum)

	cheng := cheng
	sum = operator(cheng,2,9)
	fmt.Println(sum)

	fmt.Println(operator(sub,100, 200))

	// 引用传递
	m:=8
	fmt.Println(m)
	mtest(&m)
	fmt.Println(m)

	// 给返回值命名

	j,k := calc(3,7)
	fmt.Println(j)
	fmt.Println(k)

	// 可变参数
	sum = test(10)
	fmt.Println(sum)

	sum = test(5,5,5)
	fmt.Println(sum)

	str := test2("lzl","lllll","kkkkk","nnnnn")
	fmt.Println(str)

}
// 引用传递 传递地址
func mtest(a *int)  {
	*a = 100
}

func cheng(a int,b int) int {
	return a * b
}
func operator(op add_func,a int, b int) int {
	return op(a,b)
	
}
func sub(a , b int) int {
	return a - b
}
func add(a,b int) int {
	return a + b
}

func calc(a ,b int) (sum int, avg float32) {
	sum = a + b
	avg = float32(sum) / 3
	return sum,avg
}

func test(a int,arg... int) int {
	var sum int = a
	for i:=0;i<len(arg);i++ {
		sum += arg[i]
	}
	return sum
}

func test2(a string,arg... string) (result string) {
	result = a
	for i:=0;i<len(arg);i++ {
		result += arg[i]
	}
	return result
}