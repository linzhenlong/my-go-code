package model

import "fmt"

type person struct {
	name string
	age int
	sal float64
}

func (p *person)SetSal(sal float64)()  {
	p.sal = sal
}

func (p *person)GetSal() float64  {
	return p.sal
}

func (p *person)SetAge(age int) {
	if age < 0 {
		fmt.Println("age输入有误")
	}
	if age > 150 {
		fmt.Println("age输入有误")
	}
	p.age = age
}

func (p *person)GetAge() int {
	return p.age
}

// 写一个工厂模式函数相当于构造方法
func NewPerson(name string) *person {
	return &person{
		name:name,
	}
}