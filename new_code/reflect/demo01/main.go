package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"UserName"`
	Age int `json:"age"`
} 

func main()  {
	
	var person Person
	person = Person{
		Name:"张三",
		Age:18,
	}


	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("%v\n",err)
	}
	fmt.Println(string(jsonData))
}
