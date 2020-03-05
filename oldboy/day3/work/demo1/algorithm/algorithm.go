package algorithm

import (
	"fmt"
	"strconv"
)

/**
	判断一个数是否是素数
 */
func Isprime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 水仙花数
func IsArmstrongNumber(n int) bool {
	var int_str string
	int_str = strconv.Itoa(n)
	var length int = len(int_str)
	var sum int = 0
	for i := 0; i < length; i++ {
		single_str := fmt.Sprintf("%c", int_str[i])
		single_int, _ := strconv.Atoi(single_str)
		sum += single_int * single_int * single_int
	}
	if n == sum {
		return true
	} else {
		return false
	}
}

func IsArmstrongNumber2(n int) bool {
	var(
		i int  // 个位
		j int  // 十位
		k int  // 百位
	)
	i = n % 10
	j = (n / 10) % 10
	k = (n / 100) % 100
	sum := i*i*i + j*j*j + k*k*k
	return sum==n
}
// 返回一到n的阶乘之后
func SumFactorial(n uint64) uint64 {
	var(
		res uint64
		i uint64
		j uint64
		)
	for i = 1; i <= n; i++ {
		for j = 1; j <= i; j++ {
			res = res + i*j
		}
	}
	return res
}
