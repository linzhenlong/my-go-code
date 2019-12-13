package main

import (
	"encoding/json"
	"fmt"
)
type Monster struct {
	Name string `json:"name"`  // 反序列化tag
	Age int `json:"age"`
	Birthday string `json:"Birthday"`
	Sal float64 `json:"sal"`
	Skill []string `json:"skill"`
}
// json 反序列为结构体
func json2struct()  {
	str := "{\"name\":\"张三\",\"age\":18,\"Birthday\":\"2019/10/22 21:22:47\",\"sal\":10000,\"skill\":[\"张三\",\"李四\"]}"
	var monster Monster
	err :=json.Unmarshal([]byte(str),&monster)
	if err !=nil{
		fmt.Printf("%v", err)
	}
	fmt.Println(monster)
}

func json2map() {
	str := "{\"name\":\"张三\",\"age\":18,\"Birthday\":\"2019/10/22 21:22:47\",\"sal\":10000,\"skill\":[\"张三\",\"李四\"]}"
	// 定义一个map
	mapData := map[string]interface{}{}

	// 反序列化map时不需要make,因为make操作被封装到Unmarshal()方法中
	err := json.Unmarshal([]byte(str), &mapData)
	if err !=nil{
		fmt.Printf("%v", err)
	}
	fmt.Println(mapData)
}

func json2slice()  {
	str := "[{\"address\":[\"北京\",\"上海\"],\"age\":81,\"name\":\"xxxx0\"},{\"address\":[\"北京\",\"上海\"],\"age\":87,\"name\":\"xxxx1\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str),&slice)

	if err !=nil{
		fmt.Printf("%v", err)
	}
	fmt.Println(slice)

}


func main()  {
	json2struct()
	json2map()
	json2slice()

}