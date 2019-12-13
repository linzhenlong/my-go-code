package main

import "fmt"

type Person struct {
	Name string
}

// 函数
// 接收者为值类型时,不能将指针类型的数据直接传递
func test01(person Person)  {
	fmt.Println("test01=",person.Name)
}
// 接收者为引用类型时,不能将值类型的数据直接传递
func test02(person *Person) {
	fmt.Println("test02=",(*person).Name)
}

//对于方法(如struct方法),
//接收者可以为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以
func (person Person) test03() {
	person.Name = "jack"
	fmt.Println("test03=",person.Name)
}

func (person *Person) test04() {
	(*person).Name = "jack2"
	fmt.Println("test04=",(*person).Name)
}

func main()  {
	person := Person{Name:"TOM"}
	// 调用函数
	test01(person)
	// test01(&person) 传地址会报错 cannot use &person (type *Person) as type Person in argument to test01
	// test01(&person)

	// 接收者为引用类型时,不能将值类型的数据直接传递
	test02(&person)
	// test02(person) 接收者为引用类型时,不能将值类型的数据直接传递
	// cannot use person (type Person) as type *Person in argument to test02
	//test02(person)

	// 方法
	person.test03()
	// 接收者可以为值类型时，可以直接用指针类型的变量调用方法
	fmt.Println("main person.Name",person.Name) // main person.Name TOM
	(&person).test03()

	person.test04()
	// 接收者可以为指针类型时，可以直接用值类型的变量调用方法
	fmt.Println("main person.Name test04",person.Name)
	(&person).test04() // jack2

}
