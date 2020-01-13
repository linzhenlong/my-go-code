package main

import "fmt"

type St struct {
	name string
	age int
}

func main() {
	m := make(map[string]*St)
	stus := []St{
		{
			name: "小王子",
			age: 18,
		},
		{
			name: "娜扎",
			age: 20,
		},
		{
			name: "大王八",
			age: 30,
		},
	}
	for key ,stu := range stus {
		m[stu.name] = &stus[key]
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k,"==>",v.name)
	}
}