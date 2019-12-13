package main

import "fmt"

type Person struct {
	Name string
}

func main()  {
	var p Person
	p.Name = "linzl"
	p.speak()
	p.jisuan()
	p.jisuan2(100)
	sum := p.getSum(10,20)
	fmt.Println("10+20=",sum)
}

func (person Person) speak() {
	fmt.Println(person.Name + "是一个好人！！")
}

func (person Person) jisuan() {
	sum :=0
	for i:=1;i<=1000;i++  {
		sum +=i
	}
	fmt.Println("1加到1000的和是:",sum)
}

func (p Person) jisuan2(n int)  {
	sum := 0
	for i:=0;i<=n;i++ {
		sum +=i
	}
	fmt.Printf("1加到%d的和是%d \n",n,sum)
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1+n2
}