package main

import "fmt"

//参考解析：
//这是新手常会犯的错误写法，
//for range 循环的时候会创建每个元素的副本，
// 而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
// 所以最后 map 中的所有元素的值都是变量 val 的地址，
// 因为最后 val 被赋值为3，所有输出都是3.

func main() {
	test1()
	test2()
}

func test1() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		fmt.Println(val, &val)
		m[key] = &val
	}
	fmt.Println("=====test1=======")
	for k, v := range m {
		fmt.Println(k, "-->", *v)
	}
	fmt.Println("=====test1=======")
}

func test2() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		//fmt.Println(val, &val)
		value := val
		m[key] = &value
	}
	fmt.Println("=====test2=======")
	for k, v := range m {
		fmt.Println(k, "-->", *v)
	}
	fmt.Println("=====test2=======")
}
