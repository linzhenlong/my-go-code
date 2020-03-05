package main

import (
	"fmt"
	jiecheng2 "github.com/linzhenlong/my-go-code/day2/demo18/jiecheng"
)
func main() {

	var n int = 10

	fmt.Println(jiecheng2.Add(n))


}

func jiecheng(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			res = res + i*j
		}
	}
	return res
}

func sushu(n int, m int) {
	for i := n; i <= m; i++ {

	}
}
