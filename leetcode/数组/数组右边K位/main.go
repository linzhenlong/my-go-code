package main

import "fmt"

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

// 输入:
// [1,2,3,4,5,6,7] k=3
// 输出:
// [5,6,7,1,2,3,4]

func methodA(array []int, k int) {
	length := len(array)
	if k > length {
		k = length % k // 当右移次数大于数组长度，取模后的只为新的要移动的次数
	}
	fmt.Println("before",array)
	 tempArray :=make([]int,length)
	for i:=0;i<length;i++ {
		tempArray[i] = array[i]
	}
	for i := 0; i < length; i++ {
		p := (i + k) % length
		temp := tempArray[i]
		array[p] = temp
	}
	fmt.Println("after",array)
}

func main() {
	methodA([]int{1, 2, 3, 4, 5}, 2)
}

   