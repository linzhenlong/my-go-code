package main

import "fmt"

func sum(a,b int)int  {
	return a + b
}

func sum2(n int) int  {
	res := 0
	for i:=1;i<=n;i++ {
		res +=i
	}
	return res
}


func main() {

	fmt.Print("hello world")
}
