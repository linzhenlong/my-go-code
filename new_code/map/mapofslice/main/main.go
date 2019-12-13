package main

import "fmt"

/**
要求:使用一个map来记录monster的信息name和age,也就是说一个monster对应一个map,
并且monster的个数可以动态添加 => map切片
 */
func main()  {

	//声明一个map切片
	var monsters = make([]map[string]string,1)

	// 增加一个monster
	if monsters[0] == nil {
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
		monsters[0]["sex"] = "男"
	}


	// 增加一个monster
	// 由于声明的切片的长度为1,在添加会越界
	// panic: runtime error: index out of range
	// 因此不能通过下面的方式添加

	/*if monsters[1] == nil {
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "牛魔王2"
		monsters[1]["age"] = "500"
		monsters[1]["sex"] = "男"
	}*/

	// 在添加的一个的话，需要使用到切片的append()方法,可以动态添加
	// 1先定义一个monsters map

	newMonster := make(map[string]string)
	newMonster["name"] = "新的妖怪"
	newMonster["age"] = "18"
	newMonster["sex"] = "女"
	// 2.append 进去
	monsters = append(monsters,newMonster)
	fmt.Println(monsters)

}