package main

import "fmt"

type Person struct {
	Name string
	Age int
}


func main()  {

	var p1 Person
	p1.Name = "小明"
	p1.Age = 18

	var p2 =&p1
	fmt.Println((*p2).Name)
	fmt.Println(p2.Name)
	p2.Name = "mary"
	// 由于p2 为指针结构体，p2.Name修改内容会导致p1.Name进行修改
	fmt.Printf("p1.Name=%v,p2.Name=%v \n",p1.Name,p2.Name)

}
