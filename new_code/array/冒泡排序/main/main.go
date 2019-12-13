package main

import "fmt"

func main() {

	var arr = [...]int{24, 69, 80, 57, 13}

	BubbleSort(&arr)
	fmt.Println(arr)
}

/**
  数组是数值类型，需要改变里面的内容，需要传指针进去
 */
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=",(*arr))
	for j:=0;j<len(arr)-1 ;j++  {
		for i:=0;i<len(arr)-1-j ;i++  {
			if (*arr)[i] > (*arr)[i+1] {
				temp := (*arr)[i]
				arr[i] = (*arr)[i+1]
				arr[i+1] = temp
			}
		}
	}
}
