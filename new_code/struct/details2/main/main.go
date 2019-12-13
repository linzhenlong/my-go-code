package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	Num int
}

type Student struct {
	Name string
	Age int
}

type stu Student

func main()  {

	var a A
	var b B
	fmt.Println(a,b)
	b.Num = 100
	//把b赋给a需要强转一下,不过强转需要一个前提条件
	//进行转换时需要完全相同的字段(名字，个数和类型)
	a = A(b)
	fmt.Println(a,b)

	var stu1 Student
	var stu2 stu
	stu1.Name = "宋江"
	// stu2 = stu1直接赋值会报错，需要强转，因为Golang认为是新的数据类型
	stu2 = stu(stu1)
	fmt.Println(stu1,stu2)
}


