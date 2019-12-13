package model

import "fmt"

type A struct {
	Name string
	age int
}

func (a *A) SayOk() {
	fmt.Println("A SayOK()",a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello()",a.Name)
}

type B struct {
	A
}
