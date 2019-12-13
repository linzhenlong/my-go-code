package main

import "fmt"

type Person struct {
	Name string
	Age int
}
func main()  {

	//方式1-直接声明
	var p1 Person
	p1.Name = "张三"

	//方式2-{}
	var p2 = Person{Name:"李四"}
	fmt.Println(p2)

	//方式3 - &(指针) var p3 *Person = new(Person)
	var p3 *Person = new(Person)
	// 因为p3是一个指针，因此标准的给字段赋值方式

	//这种写法 (*p3).Name = "tom"也可以直接这样写p3.Age=18
	//golang设计者为了程序员使用方便，go编译器底层对p3.Name做了转化(*p3).Name
	(*p3).Name = "tom"
	p3.Age = 18
	fmt.Println(*p3)

	// 方式4-{} var person *Person = &Person{}
	var p4 *Person = &Person{}
	// 因为p4是一个指针其标准的访问字段的方法是
	(*p4).Name = "mary"
	fmt.Println(*p4)

	//golang设计者为了程序员使用方便，也可以这样写p4.Age=10
	p4.Age = 10
	fmt.Println(*p4)
}
