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

func main() {
	fmt.Println(S)
}
