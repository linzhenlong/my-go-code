package main

import "fmt"


func main() {
	m := map[string]string{
		"name":"linzl",
		"age":"18",
		"address":"beijing",
	}
	fmt.Println(m)

	m2 := make(map[string]int) // []
	fmt.Println(m2)

	var m3 map[string]int 
	fmt.Println(m3)

	fmt.Println("Traversing map ")
	for index, v := range m {
		fmt.Println(index, v)
	}

	fmt.Println("获取map 的values")
	age,ok := m["age"]
	fmt.Printf("年龄%s,是否存在:%t\n", age, ok)
	if aeg,ok := m["aeg"]; !ok { // 故意拼错key
		fmt.Println("不存在这个key")
	} else {
		fmt.Printf("年龄%s,是否存在:%t\n", aeg, ok)
	}

	fmt.Println("删除map里的value")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
	
}