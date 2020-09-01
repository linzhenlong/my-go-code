package main

import "fmt"

func cal(n int64) int64 {
	var sum int64
	var i int64
	for i = 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func cal2(n int64) int64 {
	var sum, i, j int64
	sum = 0
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			sum = sum + i*j
		}
	}
	return sum
}

func main() {
	fmt.Printf("sum n:%d=%d\n", 10, cal(10))
	fmt.Printf("sum2 n:%d=%d\n", 10, cal2(10))
}
