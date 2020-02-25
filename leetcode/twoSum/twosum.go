package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	 for index, val := range nums {
		 for i:=index+1;i<len(nums);i++ {
			if val + nums[i] == target {
				res = append(res,index)
				res = append(res,i)
				break
			}	
		 }
	 }
	 return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 13))
}