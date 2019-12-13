package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

//Golang中的方法作用在指定的数据类型上的(即，和指定的数据类型绑定),
//因此自定义数据类型，都可以有方法，而不仅仅是struct,比如int32,
//float32等都可以有方法

type integer int

type Monster struct {
	Name string  `json:"name"`
	Age int		`json:"age"`	
	Skill string `json:"skill"`
} 

func (i integer) print() {
	fmt.Println("i=",i)
}

func (i *integer) changeI(n int) {
	*i = (*i) + integer(n)
	fmt.Printf("i加%d=%d \n",n,*i)
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.Radius,2)
}

// 为了提高效率，通常我们方法和结构的指针类型绑定
func (c *Circle) area2() float64 {

	// 因为c是指针，因此我们标准的访问其字段的方式是(*c).Radius
	// 但是由于GO底层编译器的优化，return math.Pi * math.Pow(c.Radius,2) 这样写也没有问题
	return math.Pi * math.Pow((*c).Radius,2)
}

// 如果一个变量实现了String()的方法，那么fmt.Println默认会调用这个变量的String()进行输出
func (m *Monster) String() string {
	str := fmt.Sprintf("Name = %s,Age = %d",m.Name,m.Age)
	return str
}

func main()  {

	var c Circle
	c.Radius = 4.0
	fmt.Println("面积是:",c.area())

	var d Circle
	d.Radius = 5.0

	// Go编译器底层做了优化，(&d).area2() 等价于d.area2()
	// 因为编译器会自动的加上&
	res2 := (&d).area2()
	fmt.Println("面积是",res2)

	//Golang中的方法作用在指定的数据类型上的(即，和指定的数据类型绑定),
	//因此自定义数据类型，都可以有方法，而不仅仅是struct,比如int32,
	//float32等都可以有方法
	var i integer = 100
	i.print()
	fmt.Printf("i原来=%d \n",i)
	(&i).changeI(5)
	fmt.Printf("i change后=%d \n",i)

	// 如果一个变量实现了String()的方法，那么fmt.Println默认会调用这个变量的String()进行输出
	var monster Monster
	monster.Name = "妖怪1"
	monster.Age = 500
	fmt.Println(&monster)
}
