package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
// AInterface继承BInterface,CInterface
type AInterface interface {
	BInterface
	CInterface
	test03()
}
// 如果需要实现AInterface,
// 就需要将BInterface,CInterface和本身的test03()都实现
type Stu struct {

}
func (s Stu)test01() {
	fmt.Println("test01")
}
func (s Stu)test02() {
	fmt.Println("test02")
}
func (s Stu)test03() {
	fmt.Println("test03")
}

type EmptyInterface interface {

}
func main()  {
	var stu Stu
	var a AInterface = stu
	var b BInterface = stu
	var c CInterface = stu
	b.test01()
	c.test02()
	a.test03()
	a.test01()
	var empty EmptyInterface = stu
	fmt.Println(empty) // 输出{}
	var empty2 interface{} = stu
	fmt.Println(empty2) // 输出{}
	var i int
	var empty3 EmptyInterface = i
	fmt.Println(empty3)
}