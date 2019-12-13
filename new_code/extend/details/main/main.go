package main

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

func (b *B) hello() {
	fmt.Println("b hello()",b.Name)
}

type B struct {
	A
	Name string
}

func main()  {

	/*var b B
	b.A.Name = "TOM"
	b.A.age = 19

	b.A.SayOk()
	b.A.hello()*/

	var b B
	b.Name = "jack"
	b.age = 100
	b.SayOk() // A SayOK()
	b.hello() // b hello() jack

	b.A.Name = "TOM"
	b.SayOk() // A SayOK() TOM
	b.hello() // b hello() jack
	b.A.hello() //A hello() TOM
}