package main

import "github.com/linzhenlong/my-go-code/2020/data_structure/动态数组/001_demo/002_demo/array"

import "fmt"

func main() {
	arr := array.NewArray(10)

	for i := 0; i <= 5; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)
	arr.Add(4, 10)
	fmt.Println(arr)
	arr.AddFirst(999)
	fmt.Println(arr)
}
