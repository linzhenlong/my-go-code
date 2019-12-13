package main

import (
	"fmt"
	"sort"
)

func main()  {
	var nums = [...]int{41,9,101,127,20,999,2000,8823,10,88,66}
	fmt.Println(nums)
	sort.Ints(nums[:])
	fmt.Println(nums)

	fmt.Println("请输入一个数。。。")
	var findNum int
	fmt.Scanln(&findNum)
	// sort.SearchInts 查找
	index := sort.SearchInts(nums[:],findNum)
	fmt.Println(index)

	// 二分法查找
	BinaryFind(&nums,0,len(nums)-1,findNum)
}

func BinaryFind(arr *[11]int, leftIndex int, rightIndex int, findValue int) {

	// 跳出递归条件
	if leftIndex > rightIndex {
		fmt.Println("找不到了")
		return
	}
	// 找到中间的下标
	middleIndex := (rightIndex + leftIndex)/2
	fmt.Printf("middleindex=%d\n",middleIndex)
	if (*arr)[middleIndex] > findValue {
		BinaryFind(arr,leftIndex,middleIndex-1,findValue)
	} else if (*arr)[middleIndex] < findValue {
		BinaryFind(arr,middleIndex+1,rightIndex,findValue)
	} else {
		fmt.Printf("找到了，下标是%v",middleIndex)
	}
}