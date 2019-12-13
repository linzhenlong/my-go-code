package main

import (
	"fmt"
	"sort"
)

func main()  {
	map1 := make(map[int]int,100)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 77

	fmt.Println(map1)

	// 如果按照map的key的顺序进行排序输出
	// 1.先将map的key放到一个切片中
	var slice1 []int
	for i,_ := range map1{
		slice1 = append(slice1,i)
	}
	fmt.Println(slice1)

	// 2.对切片进行排序
	sort.Ints(slice1)

	// 3.遍历切片，然后按照key 来输出
	for _,v := range slice1 {
		fmt.Println(map1[v])
	}

}
