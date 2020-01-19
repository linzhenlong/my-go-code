package main

import "fmt"

// 数组是值类型
func printArray(arr [5]int) {
	arr[0] = 1000
	for _, v := range arr {
		fmt.Println(v)
	}
}

func printArrayPointer(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i,v)
	}
}

func main() {
	// 定义数组
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}

	// 二维数组
	var grid [4][5]int // 4行5列

	grid2 := [3][2]int{
		{
			2,3,
		},
		{
			4,6,
		},
		{
			7,8,
		},
	}
	fmt.Println(
		arr1,
		arr2,
		arr3,
		grid,
		grid2,
	)

	// 遍历
	for i:=0;i<len(arr3);i++ {
		fmt.Println(arr3[i])
	}
	for i:= range arr3 {
		fmt.Println(arr3[i])
	}
	for i,v := range arr3 {
		fmt.Println(i, v)
	}

	printArray(arr1)
	printArray(arr3)
	//printArray(arr2) // cannot use arr2 (type [3]int) as type [5]int in argument to printArray
	fmt.Println(arr1, arr3)

	printArrayPointer(&arr1)
	printArrayPointer(&arr3)
	fmt.Println(arr1,arr3)
}
