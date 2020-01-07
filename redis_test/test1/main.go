package main

import "fmt"

func main() {
	var  i,a int
	for a < 5 {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		fmt.Println("for 循环外", i)
		a++
	}
}
