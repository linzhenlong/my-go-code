package main

import "fmt"

func main()  {
	a()
	fmt.Println("###############")
	f()
	fmt.Println(test(100,200))
}

func a()  {
	i:=15
	defer fmt.Println(i)
	defer fmt.Println("2222")
	i++
	fmt.Println(i)
	return
}

func f()  {
	for i:=0;i<6;i++ {
		defer fmt.Println(i)
	}
}
func test(a , b int) int {
	result := func(a1, b1 int) int {
		return a1 + b1
	}
	return result(a, b)
}