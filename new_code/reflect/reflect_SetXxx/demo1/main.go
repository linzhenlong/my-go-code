package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func TestSetInt(b interface{})  {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal 的kind:%v\n",rVal.Kind())
	// 因为b 传进来是一个指针类型,因此获取及改变值的时候需要使用Elem()方法
	fmt.Printf("修改前的值%v\n",rVal.Elem())
	rVal.Elem().SetInt(1000)
	fmt.Printf("修改后的值%v\n",rVal.Elem())
}

// 修改结构体的值
func TestSetStructField(b interface{})  {
	rVal := reflect.ValueOf(b)
	fileName :=rVal.Elem().FieldByName("Name")
	fileAge := rVal.Elem().FieldByName("Age");

	fmt.Printf("修改前%s\n",fileName)
	fileName.SetString("王五")
	fileAge.SetInt(20)
	fmt.Printf("修改后%s\n",fileName)
}

func TestFloat(b interface{})  {
	rVal := reflect.ValueOf(b)
	rType := reflect.TypeOf(b)
	fmt.Printf("b的Type是:%v\n",rType)
	fmt.Printf("b的Kind是:%v\n",rVal.Kind())
	fmt.Printf("b的Kind是:%v\n",rType.Kind())

	// 转接口
	iV := rVal.Interface()
	fmt.Printf("b的值%f,类型%T\n",iV.(float64), iV.(float64))

}

func main() {

	var num int = 100
	TestSetInt(&num)
	fmt.Printf("num=%d\n",num)
	person := Person{
		Name:"张三",
		Age:18,
	}
	fmt.Printf("%v\n",person )

	TestSetStructField(&person)
	fmt.Printf("%v\n",person )

	var num2 = 3.1415926
	TestFloat(num2)

	var str string = "TOM"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Printf("%s\n",str)

	var num3 int = 200
	var ptr *int = &num3
	fmt.Println(*ptr,num3)
	*ptr = 500
	fmt.Println(*ptr,num3)

}
