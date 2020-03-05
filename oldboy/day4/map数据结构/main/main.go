package main

import "fmt"

func main()  {
	testMap()
	testMap2()
	testMap3()
	testMap4()

	var a = make(map[string]int)
	a["one"] = 678
	modify1(a)
	fmt.Println(a) // map 为引用数据类型 ,在函数中修改，原值也被修改

	testSliceOfMap()
}

func testMap()  {

	// make 声明初始化
	var mapA = make(map[string]string,5)

	mapA["hhh"] = "ooo"
	fmt.Println(mapA)

	//直接声明并初始化
	var mapB = map[string]string{
		"name" : "linzl",
		"age":"20",
	}
	fmt.Println(mapB)

	var mapC = map[string]string{"hello":"name"}
	fmt.Println(mapC)
}

func testMap2()  {
	var mapA = make(map[string]map[string]string,5)
	mapA["key1"] = make(map[string]string)
	mapA["key1"]["name1"] = "王五"
	mapA["key1"]["name2"] = "王五"
	mapA["key1"]["name3"] = "王五"
	mapA["key1"]["name4"] = "王五"

	fmt.Println(mapA)
}

func testMap3()  {
	var a = map[string]string{"hello":"world"}
	fmt.Println(a)

	a = make(map[string]string,10)

	a["hello"] = "时间"
	fmt.Println(a)

	// 查找
	val, ok := a["hello"]
	fmt.Println(val, ok)

	a["hello1"] = "test"
	a["hello2"] = "test2"
	a["hello3"] = "test3"

	for key, val := range a {
		fmt.Println(key,val)
	}
	delete(a,"hello3")

	fmt.Println(len(a))
}

func modify(a map[string]map[string]string,key string)  {
	_, ok := a[key]

	// 如果不存在 初始化
	if !ok {
		a[key] = make(map[string]string)
	}
	a[key]["password"] = "12345"
	a[key]["nickname"] = "小三"
}

func testMap4()  {
	var a = make(map[string]map[string]string)
	a["linzl"] = make(map[string]string)
	a["linzl"]["password"] = "123"
	modify(a,"linzl")

	fmt.Println(a)
}

func modify1(a map[string]int)  {
	a["one"] = 234
}

func testSliceOfMap()  {
	var sliceMap = make([]map[int]int, 5)
	for i:=0;i<2;i++ {
		sliceMap[i] = make(map[int]int)
		sliceMap[i][i] = i
	}
	fmt.Println(sliceMap)
}