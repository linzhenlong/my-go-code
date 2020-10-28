package main

import "fmt"

//Parent ...
type Parent struct {

}
// MethodA ...
func (p *Parent)MethodA() {
	fmt.Println("MethodA from Parent")
	p.MethodB()
}
// MethodB ...
func (p *Parent)MethodB() {
	fmt.Println("MethodB from Parent")
}
// Child ...
type Child struct {
	Parent
}
// MethodB ...
func (b *Child)MethodB() {
	fmt.Println("MethodB from child")
}

func main() {
	child := Child{}
	child.MethodA()
	child.MethodB()
}