package main

import "fmt"

func main()  {

	var arr = [...][4]int{{1,2,3,45},{1,2,3,46}}

	// 双层for循环遍历
	for i:=0;i<len(arr);i++ {
		for j:=0;j<len(arr[i]);j++ {
			fmt.Printf("arr[%d][%d]=%d\t",i,j,arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	// for range
	for i,v := range arr {
		for j,val2 := range v {
			fmt.Printf("arr[%d][%d]=%d\t",i,j,val2)
		}
		fmt.Println()
	}
}
