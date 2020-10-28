package main

import "fmt"

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

func (d Direction) String() string {
	return []string{"N", "E", "S", "W"}[d]
}

func test1() {
	vals := []int{0, 1, 2, 3}
	valptr := make([]*int, 0)
	for _, v := range vals {
		valptr = append(valptr, &v)
	}
	//fmt.Println(valptr)
	for _,v := range valptr {
		fmt.Println(*v)
	}
}

func main() {
	//fmt.Println(S)
	test1()
}
