package main

import "encoding/json"

import "fmt"

import "reflect"

// 下面代码输出什么 
// 输出float64 
// 因为在json.Unmarshal() 源码中有写:

// To unmarshal JSON into an interface value,
// Unmarshal stores one of these in the interface value:
//
//	bool, for JSON booleans
//	float64, for JSON numbers
//	string, for JSON strings
//	[]interface{}, for JSON arrays
//	map[string]interface{}, for JSON objects
//	nil for JSON null

func main() {
	jsonStr := []byte(`{"age":1}`)
	var value map[string]interface{}
	json.Unmarshal(jsonStr,&value)
	age := value["age"]
	fmt.Println(reflect.TypeOf(age)) 
	fmt.Println(reflect.ValueOf(age)) 
	fmt.Println(reflect.ValueOf(age).Kind()) 
}