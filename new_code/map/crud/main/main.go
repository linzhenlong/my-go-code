package main

import "fmt"

func main()  {

	cities :=make(map[string]string)
	cities["c1"] = "北京"
	cities["c2"] = "上海"
	cities["c3"] = "广州"
	fmt.Println(cities)

	cities["c3"] = "广州new" // 更新
	fmt.Println(cities)

	// 删除
	delete(cities,"c3")
	delete(cities,"c8") //删除一个不存在的key 也不报错
	fmt.Println(cities)

	// 查找
	val,ok := cities["c1"]
	if ok {
		fmt.Println("找到了c1对应的value了",val)
	} else {
		fmt.Println("木有找到了c1对应的value")
	}

}
