package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)
const TimeFormat  = "2006/01/02 15:04:05"
// 定义一个结构体

type Monster struct {
	Name string `json:"name"`  // 反序列化tag
	Age int `json:"age"`
	Birthday string `json:"Birthday"`
	Sal float64 `json:"sal"`
	Skill []string `json:"skill"`
}

func struct2json()  {
	monster := Monster{
		Name:"张三",
		Age:18,
		Birthday:time.Now().Format(TimeFormat),
		Sal:10000.00,
		Skill:[]string{"张三","李四"},
	}
	jsonStr, err :=json.Marshal(monster)
	if err !=nil {
		fmt.Println(nil)
	}
	fmt.Println(string(jsonStr))
}

func map2json()  {

	// 定义一个map.
	//resp := map[string]interface{} {}
	var resp map[string]interface{}
	resp = make(map[string]interface{})
	resp["error_code"] = 0
	resp["msg"] = "success"


	var row []Monster
	//row := make([]Monster, 10)
	for i:=0; i<4;i++ {
		monster :=  Monster{
			Name:"张三",
			Age:18+i,
			Birthday:time.Now().Format(TimeFormat),
			Sal:10000.00,
			Skill:[]string{"张三","李四"},
		}
		row = append(row, monster)
	}

	data :=map[string]interface{} {}
	data["row"] = row
	data["total"] = 10
	resp["data"] = data
	dataJson , err := json.Marshal(resp)
	if err != nil {
		fmt.Println(nil)
	}
	fmt.Println(string(dataJson))
}

func slice2json()  {
	var slice []map[string]interface{}
	//slice = make([]map[string]interface{}, 5)

	for i:=0;i<2;i++{
		m := make(map[string]interface{})
		m["name"] = "xxxx"+ strconv.Itoa(i)
		m["age"] = rand.Intn(100)
		m["address"] = []string{"北京","上海"}
		slice = append(slice,m)
	}
	dataJson , err := json.Marshal(slice)
	if err != nil {
		fmt.Println(nil)
	}
	fmt.Println(string(dataJson))

}
func main()  {
	//将结构体，map,切片进行序列化
	struct2json();
	map2json();
	slice2json()
}

