package main

import "fmt"

const (
	x = iota //0
	_
	y //2
	z = "zz"
	k //zz 
	p = iota 
)

func main() {
	fmt.Println(x, y, z, k, p)
}
