package main

import "fmt"

/**
 输出
	000000
	000100
	020300
	000000
 */
func main()  {
	var arr [4][6]int

	arr[1][3] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	for i:=0;i<len(arr);i++ {
		for j:=0;j<len(arr[i]);j++ {
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}

	var arr2 [2][3]int // 分析内存布局

	arr2[1][1] = 10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址:%p\n",&arr2[0])
	fmt.Printf("arr2[1]的地址:%p\n",&arr2[1])
	fmt.Printf("arr2[0][0]的地址:%p\n",&arr2[0][0])
	fmt.Printf("arr2[1][0]的地址:%p\n",&arr2[1][0])

	// 另外一种赋值方式
	var arr3  = [2][3]int{{1,2,3},{2,3,4}}
	fmt.Println(arr3)


	var arr4 = [...][3]int{{1,2,3},{2,3,4}}
	fmt.Println(arr4)

}