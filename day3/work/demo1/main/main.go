package main

import (
	"fmt"
	"go_dev/day3/work/demo1/algorithm"
)

func main() {

	//判断素数 输出101~200之间的素数并记录个数
	fmt.Println("========素数开始========")
	var count int = 0
	for i:=101;i<200;i++ {
		if algorithm.Isprime(i) {
			count++
			fmt.Println(i)
		}
	}
	fmt.Println("total:",count)
	fmt.Println("========素数结束========")
	fmt.Println("=======================")

	fmt.Println("========水仙花数开始========")

	for i:=100;i<1000 ;i++  {
		if algorithm.IsArmstrongNumber2(i) {
			fmt.Println(i)
		}
	}
	fmt.Println("========水仙花数结束========")
	fmt.Println("=======================")
	fmt.Println("========1到n的阶乘========")
	fmt.Println(algorithm.SumFactorial(10))


}


