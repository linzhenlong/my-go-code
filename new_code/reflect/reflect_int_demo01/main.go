package main

import (
	"fmt"
	"reflect"
)

// 专门演示反射的
func reflectTest01(b interface{}) {
	// 通过反射获取到传入变量的，type(类型) ，kind(类别)
	// 1.先获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("int num 的reflect.TypeOf():%s\n",rType)

	// 2.获取reflect.Value类型
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v,type=%T\n",rValue, rValue)
	fmt.Printf("rValue.Kind()=%v\n",rValue.Kind())


	n1 := 10
	n2 := n1 + (int)(rValue.Int()) // 转为原始的int
	fmt.Println(n2)

	// 将rValue 转成interface{}类型
	iV := rValue.Interface()

	// interface通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println(num2)
}


func main()  {

	// 定义一个基本类型变量
	var num int = 10
	reflectTest01(num)
}

