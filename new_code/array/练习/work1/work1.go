package work1

import (
	"fmt"
	"go_dev/new_code/array/练习/sort"
	"math/rand"
	"time"
)

/*
  随机生成10个数，0-100范围，保存到数组，并倒序打印，以及求平均值，及最大值的下标，并查找里面是否有55
 */
func WorkOne() (arr []int,avg float64,bigNUmIndex int,isHas bool){

	rand.Seed(time.Now().UnixNano())
	var nums [10]int
	for i:=0;i<10;i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println("排序前",nums);
	arr = sort.BubbleSort(nums[:],"asc")


	sum :=0
	big :=0
	for _,v := range nums {
		sum +=v
		if v > big {
			big = v
		}
		if v == 55 {
			isHas = true
		}
	}

	for i,v := range nums {
		if v == big {
			bigNUmIndex = i
		}
	}
	avg = float64(sum / len(arr))

	fmt.Println("排序后",arr);
	fmt.Println("平均值",avg);
	fmt.Println("最大值索引",bigNUmIndex);
	fmt.Println("是否含有55",isHas);
	return
}
