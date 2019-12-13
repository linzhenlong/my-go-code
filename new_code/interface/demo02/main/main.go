package main

import "fmt"


type Student struct {
	Name string
}
func (s *Student) Say()  {
	fmt.Println("Student Say()")
}

type integer int
type BInterFace interface {
	Sum(a integer,b integer) integer
}
type AInterFace interface {
	Say()
}
func (i integer)Sum(a integer,b integer) integer  {
	return a + b
}
func (i integer) Say()  {
	fmt.Println("integer Say()")
}


func main()  {
	// 结构体变量,并实现了接口里的所有方法
	/*var stu Student
	// 结构体变量赋值给接口实例
	var a AInterFace = &stu
	a.Say()*/
	// 自定义类型 integer 同时实现AInterFace与BInterFace
	var i integer = 10
	var a AInterFace = i
	var b BInterFace = i
	fmt.Println(b.Sum(i,i)) // 输出20
	a.Say()

}
