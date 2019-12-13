package main

import "fmt"

// 定义一个结构体
type student struct {
	Name string
	Age int
	Grade float64
}

func main()  {
	// map是引用类型，遵守引用类型值传递的机制，
	// 在一个函数接收map,修改后，会直接修改原来的map

	var map1 = map[string]string{
		"no1" : "宋江",
	}
	fmt.Println("修改前",map1)
	modify(map1)
	fmt.Println("修改后",map1)


	//map的容量达到后，在想map添加元素，会自动扩容，
	// 并不会发生panic,也就是说map能动态的增长键值对(key-value)

	var map2 = make(map[int]int,1)
	map2[1] = 100
	map2[2] = 100
	map2[3] = 100
	map2[4] = 100
	fmt.Println(map2)

	// 与切片的区别 切片自动扩容只能通过append
	var slice1 = make([]int,1)
	slice1[0] = 100
	// slice1[2] = 100 报panic: runtime error: index out of range
	slice1 = append(slice1,200)
	fmt.Println(slice1)


	//map的value也经常使用struct类型，
	// 更适合管理复杂的数据（比之前value是一个map更好），
	// 比如value为student结构体

	// 1.map的可以是学号是唯一的
	// 2.map的value 是结构体 struct,包含学生的名称，年龄，地址
	
	students := make(map[int]student)

	// 创建学生
	stu1 :=student{"张三",18,90.5}
	students[1] = stu1

	stu2 :=student{"李四",28,95.5}
	students[2] = stu2

	stu3 := student{Name:"王五",Age:34,Grade:100.0}
	students[3] = stu3
	fmt.Println(students)
	fmt.Println(stu3.Name)

	// 遍历
	for i,v :=range students{
		fmt.Printf("学号:%d,姓名:%s,年龄:%d,成绩:%f\n",i,v.Name,v.Age,v.Grade)
	}



}

func modify(map1 map[string]string)  {
	map1["no1"] = "modify修改"
}