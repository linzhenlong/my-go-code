package main

import "fmt"

type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int // 指针  引用类型
	slice []int // 切片  引用类型
	map1 map[string]string // map 引用类型

}

type Monster struct {
	Name string
	Age int
}


func main()  {

	//指针，slice,和map的零值都是nil,即还没有分配空间，
	// 如果使用这样的字段需要先make才能使用

	// 定义一个结构体变量

	var person Person

	fmt.Println(person)  // { 0 [0 0 0 0 0] <nil> [] map[]}

	if person.ptr == nil {
		fmt.Println("ptr还没有分配空间")
	}

	if person.slice == nil {
		fmt.Println("slice还没有分配空间")
	}

	if  person.map1 == nil {
		fmt.Println("map1还没有分配空间")
	}

	// person.slice[0] = 100 //还没有赋值直接分配空间会报错
	// fmt.Println(person.slice)

	// 切片是引用类型，需要先make
	person.slice = make([]int,2)
	person.slice[0] = 100
	person.slice = append(person.slice,100)
	fmt.Println(person.slice)

	// 使用map,map是引用类型，需要先make
	person.map1 = make(map[string]string)
	person.map1["name"] = "张三"
	person.map1["age"] = "18岁"
	fmt.Println(person.map1)

	person.ptr = new(int)
	num := 10086
	person.ptr = &num
	fmt.Println(person.ptr)

	//不同结构体变量的字段是独立，互不影响，
	//一个结构体变量字段的更改，不影响另外一个,结构体是值类型

	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1

	fmt.Printf("monster1的地址%p,值是:%v \n",&monster1,monster1)
	fmt.Printf("monster2的地址%p,值是:%v \n",&monster2,monster2)
	monster2.Name = "牛魔王@@@"
	fmt.Printf("monster1的地址%p,值是:%v \n",&monster1,monster1)
	fmt.Printf("monster2的地址%p,值是:%v \n",&monster2,monster2)


}
