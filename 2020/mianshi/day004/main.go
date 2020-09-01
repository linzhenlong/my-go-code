package main

import "fmt"

// 观察下面代码是否可以编译通过.
func test1() {
	//list := new([]int) //不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，
	list := make([]int, 0)
	list = append(list, []int{1, 2, 3}...) // 不能对指针执行 append 操作。可以使用 make() 初始化之后再用。
	fmt.Println("test1", list)
	// 同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。
}

func test2() {
	s1 := []int{1, 3, 4, 5}
	s2 := []int{6, 7, 8, 9}
	//s1 = append(s1, s2)
	//不能通过编译。append() 的第二个参数不能直接使用 slice，需使用 … 操作符，
	// 将一个切片追加到另一个切片上：append(s1,s2…)。或者直接跟上元素，形如：append(s1,1,2,3)
	s1 = append(s1, s2...)
	size := 1000 // 函数内部使用简短模式.
	fmt.Println("test2", s1, size)
}

var (
	// size :=1000  // 不能通过编译
	/**
	1.必须使用显示初始化；
	2.不能提供数据类型，编译器会自动推导；
	3.只能在函数内部使用简短模式
	**/

	size    int = 1000
	maxSize     = size * 2
)

func main() {
	test1()
	test2()
	fmt.Println(size, maxSize)
}
