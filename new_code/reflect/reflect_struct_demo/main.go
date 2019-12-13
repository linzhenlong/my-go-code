package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	age int
	Score float64
}

func TestReflectStruct(b interface{})  {
	// 通过反射获取到传入变量的，type(类型) ，kind(类别)
	// 1.先获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("struct student的type是:%s\n",rType)
	fmt.Printf("struct student的Kind()是:%s\n", rType.Kind())

	// 2.获取reflect.Value类型
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v,type=%T\n",rValue, rValue)
	fmt.Printf("rValue.Kind()=%v\n",rValue.Kind())

	// 将rValue 转成interface
	iV := rValue.Interface()

	fmt.Printf("iV=%v\n",iV)

	// 接口类型断言
	// 最好判断一下类型
	stu, ok:= iV.(Student)
	if ok {
		name := stu.Name
		fmt.Printf("Student.Name=%s\n",name)
		fmt.Printf("Student.age=%d\n",stu.age)
	}


}

func main()  {

	// 定义一个Student 结构体实例
	student := Student{
		Name:"李四",
		age:18,
		Score:88.0,
	}
	TestReflectStruct(student)
}
