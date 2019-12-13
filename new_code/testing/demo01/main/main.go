package main

import "fmt"
// 被测试的函数
func addUpper(n int) int {
	res := 0
	for i:=1;i<=n;i++ {
		res +=i
	}
	return res
}

func main() {

	// 传统测试方法，直接在main方法中测试使用，看看结果是否正常
	res := addUpper(10) // 从1+到10=55
	if res !=55 {
		fmt.Printf("addUpper()有误 返回值%d期望值%d\n", res,55)
	} else {
		fmt.Printf("addUpper()正确 返回值%d期望值%d\n", res,55)
	}

	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%v len=%d cap=%d 内存:%p %v\n",
		s, len(x), cap(x), &x, x)
}
