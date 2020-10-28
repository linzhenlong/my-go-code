package main

import "fmt"

/**
 * https://leetcode-cn.com/problems/running-sum-of-1d-array/submissions/
 */

func runningSum(nums []int) []int {
	var res = make([]int,len(nums))
	for i := 0; i < len(nums); i++ {
		var sum = 0
		for j := 0; j <= i; j++ {
			sum += nums[j]
		}
		res[i] = sum
	}
	return res
}
func main() {
	nums := []int{3,1,2,10,1}
	fmt.Println(runningSum(nums))
}
