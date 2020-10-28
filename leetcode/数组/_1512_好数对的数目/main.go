package main

import "fmt"

/**
https://leetcode-cn.com/problems/number-of-good-pairs/
**/

func numIdenticalPairs(nums []int) int {
	count := 0
	for i:=0;i<len(nums);i++ {
		for j:=0;j<=i;j++{
			if nums[i] == nums[j] {
				//if i<j{
					count++
					continue
				//}
			}
		}
	}
	return count
}
func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums))
}
