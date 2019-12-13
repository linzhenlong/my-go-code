package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"name"` // `json:"name"` 结构体标签, 转json时字段名小写
	Age int	`json:"age"`
	Skill string `json:"skill"`
}

func main()  {

	//1.创建Monster变量
	var monster Monster
	monster.Name = "牛魔王"
	monster.Age = 500
	monster.Skill = "狮吼功"

	//2.将monster序列化为json

	monsterJson,err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(monsterJson)) // {"name":"牛魔王","age":500,"skill":"狮吼功"}

}