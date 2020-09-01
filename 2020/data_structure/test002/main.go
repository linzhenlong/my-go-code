package main

import "fmt"

func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	if n <= 1 {
		return n
	}
	var (
		first  int = 0
		second int = 1
	)
	for i := 0; i < n-1; i++ {
		sum := first + second
		first = second
		second = sum
	}
	return second
}

func main() {
	fmt.Println(fib1(1))
	fmt.Println(fib1(2))
	fmt.Println(fib1(3))
	fmt.Println(fib1(4))
	fmt.Println(fib1(5))
	fmt.Println(fib1(6))
	fmt.Println(fib1(7))
	fmt.Println("##################")
	fmt.Println(fib2(1))
	fmt.Println(fib2(2))
	fmt.Println(fib2(3))
	fmt.Println(fib2(4))
	fmt.Println(fib2(5))
	fmt.Println(fib2(6))
	fmt.Println(fib2(7))
	fmt.Println(fib2(64))
}
