package main

import (
	"fmt"
)

func add() func(int) int {
	var x = 10
	return func(d int) int {
		x += d
		return x
	}
}

func main() {
	f := add()
	fmt.Println(f(10)) // 20
	fmt.Println(f(10)) // 30 ,因为此时add 里的x变为20了

	f2 := add()
	fmt.Println(f2(5))
	fmt.Println(f2(10))
}
