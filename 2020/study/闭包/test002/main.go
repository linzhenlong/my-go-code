package main

import "fmt"

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func main() {
	f := add(10)
	fmt.Println(f(1), f(2))
}
