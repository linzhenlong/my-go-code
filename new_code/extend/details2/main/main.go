package main

import "fmt"

type A struct {
	Name string
	Age int
}

type B struct {
	Name string
	Score float64
}

type C struct {
	A
	B
}

func (a *A)Hello()  {
	fmt.Println("A hello")
}
func (b *B)Hello()  {
	fmt.Println("B hello")
}

type D struct {
	a A  // 有名结构体
}

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main()  {

	var c C

	// 如果C 没有Name字段，则必须通过指定匿名结构体名来区分，否则会报错
	//c.Name = "TOM" // 这种方式会报错，编译器不知道为那个结构体去赋值了，ambiguous selector c.Name
	c.A.Name = "A-tom"
	fmt.Printf("c.A.Name=%s\n",c.A.Name)
	fmt.Printf("c.B.Name=%s\n",c.B.Name)
	c.B.Name = "B-TOM"
	fmt.Printf("c.B.Name=%s\n",c.B.Name)
	c.Age = 100
	c.Score = 90.0
	fmt.Println(c)

	var d D
	// 以下这种写法会报错,因为嵌套结构体A 是一个有名结构体，
	// 需要这种方式使用d.a.xxx 里的字段或是方法
	// d.Name = "臭小子"
	d.a.Name = "臭小子"
	fmt.Println(d.a)
	fmt.Println(d)



	tv := TV{
		Goods{
			Name:"电视机001",
			Price:5999,
		},
		Brand{
			Name:"松下",
			Address:"日本",
		},
	}
	fmt.Println(tv)

	tv2 := TV2 {
		&Goods{
			Name:"电视机002",
			Price:1999,
		},
		&Brand{
			Name:"海尔",
			Address:"青岛",
		},
	}
	fmt.Println(*tv2.Goods,*tv2.Brand)
}