package main

import "fmt"

// Monkey 结构体
type Monkey struct {
	Name string
}


type LittleMonkey struct {
	Monkey //基础
}

// 声明接口
type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

// 让littleMonkey 实现BirdAble接口的Flying方法
func (m LittleMonkey)Flying() {
	fmt.Printf("'%s'可以像鸟一样飞翔了...\n",m.Name)
}
// 让littleMonkey 实现FishAble接口的Swimming方法
func (m LittleMonkey)Swimming() {
	fmt.Printf("'%s'可以像鱼儿一样游泳...\n",m.Name)
}


func (m *Monkey)Run() {
	fmt.Printf("'%s'在跑步...\n",m.Name)
}
func main()  {

	// 定义一个LittleMonkey实例
	little_monkey := LittleMonkey{
		Monkey{
			Name:"小猴子",
		},
	}
	little_monkey.Run()
	little_monkey.Flying()
	little_monkey.Swimming()
}
