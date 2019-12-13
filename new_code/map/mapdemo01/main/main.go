package main

import "fmt"

func main()  {
	// map的声明和注意事项

	var a map[string]string
	//a["no1"] = "宋江" // map 没有初始化直接赋值会报错 panic: assignment to entry in nil map

	// 在使用map 前需要先make,make的作用就是给map分配内存空间
	a = make(map[string]string,5)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松" // 覆盖，因此key不可以重复，重复会被覆盖
	a["no3"] = "武松" // 不覆盖，value可以重复
	fmt.Println(a)


	// 声明，这是map=nil
	var names map[string]string
	// make 分配一个空map
	names = make(map[string]string,10)
	fmt.Println(names)

	// 声明直接赋值
	var map2 = map[string]string{
		"no1":"武松",
	}
	fmt.Println(map2)
}
