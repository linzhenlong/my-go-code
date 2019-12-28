package main

import "fmt"

func sum(n int) int  {
	res := 0
	for i:=1;i<=n;i++ {
		res +=i
	}
	return res
}

func main()  {

	res := sum(10)
	fmt.Print(res)
}