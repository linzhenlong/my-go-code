package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type User struct {
	Name string `json:"userName"`
	Age int `json:"age"`
	Birthday string `json:"birthday"`
	SKill []string
}

// 结构体转json
func structToJson() (string ,error) {
	var user User
	skill := []string{"学习","调完"}
	user = User{
		Name:"张三",
		Age:18,
		Birthday:time.Now().Format("2006-01-02 15:04:05"),
		SKill:skill,
	}
	data , err := json.Marshal(user)
	return string(data),err
}

// json 转结构体.
func jsonToStruct()  {
	str ,err := structToJson()
	if err != nil {
		fmt.Printf("struct to json err:%v\n",err)
	}
	var user User
	err2 := json.Unmarshal([]byte(str), &user)
	if err2 != nil {
		fmt.Printf(" json to struct err:%v\n",err2)
	}
	fmt.Printf("josn:%v to struct:%v\n", str, user)
}

// map to json
func mapToJson() (string, error) {
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["error_code"] = 0
	data["error_msg"] = "";
	var row []User
	for i:=0;i<3;i++  {
		var user User
		user = User{
			Name:"user"+strconv.Itoa(i),
			Age:18,
		}
		row = append(row,user)
	}
	data["data"] = row
	jsonStr , err := json.Marshal(data)
	return string(jsonStr),err
}

// json to map
func jsonToMap() {
	str ,err := mapToJson()
	if err != nil {
		fmt.Printf("map to json err:%v\n",err)
	}
	var data map[string]interface{}
	// json 反序列化map 时 不需要make 申请内存空间，因为make操作已经被封装到Unmarshal
	err2 :=json.Unmarshal([]byte(str),&data)
	if err2 != nil {
		fmt.Printf(" json to map err:%v\n",err2)
	}
	fmt.Printf("josn:%v to map:%v\n", str, data)
}

func sliceToJson()([]byte, error) {
	var data []User
	var user User
	user.Name = "张三";
	user.Age = 18
	data = append(data, user)
	return json.Marshal(data)

}

func jsonToSlice()  {
	str, err := sliceToJson()
	if err != nil {
		fmt.Printf("slice to json err:%v\n",err)
	}
	var slice []User
	err2 := json.Unmarshal(str,&slice)
	if err2 != nil {
		fmt.Printf(" json to slice err:%v\n",err2)
	}
	fmt.Printf("josn:%v to slice:%v\n", string(str), slice)
}

func main()  {
	fmt.Println("结构体与json互相转换:")
	jsonToStruct()
	fmt.Println("map与json互相转换:")
	jsonToMap()
	fmt.Println("slice与json互相转换:")
	jsonToSlice()
}
