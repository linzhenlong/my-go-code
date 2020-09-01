package main

import "fmt"

type myInt1 int
type myInt2 int

func main() {
	var i int = 100
	//var i1 myInt1 = i //cannot use i (type int) as type myInt1 in assignment
	//var i2 myInt1 = i //cannot use i (type int) as type myInt1 in assignment
	var i1 myInt1 = myInt1(i)
	var i2 myInt2 = myInt2(i)
	fmt.Println(i1, i2)
}
