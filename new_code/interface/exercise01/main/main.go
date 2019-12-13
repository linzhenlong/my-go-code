package main

import "fmt"

type AInterFace interface {
	Test01()
	Test02()
}
type BInterFace interface {
	Test02()
	Test03()
}
type Stu struct {

}

func (s Stu)Test01() {

}
func (s Stu)Test02() {

}
func (s Stu)Test03() {

}
func main()  {
	stu := Stu{}
	var a AInterFace = stu
	var b BInterFace = stu
	fmt.Println("ok ~",a,b)

}
