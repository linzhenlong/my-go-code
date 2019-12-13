package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

// 专门用作反射
func TestReflect(val interface{})  {

	// 将interface 转为reflect.Value
	reflectVal := reflect.ValueOf(val)
	fmt.Println(reflectVal.Type())

	// 在将reflect.Value 转回interface{}
	interfaceVal := reflectVal.Interface()
	fmt.Println(interfaceVal)

	// 如果将interface 原来的变量类型，使用类型断言
	v := interfaceVal.(Student)
	fmt.Println(v)

}

func main()  {

	student := Student{
		Name:"小米",
		Age:18,
	}
	TestReflect(student)
}
