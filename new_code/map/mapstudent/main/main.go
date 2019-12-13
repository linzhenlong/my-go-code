package main

import "fmt"


//
func main()  {
	var students = make(map[int]map[string]string)

	for i:=1;i<5;i++  {
		students[i] = make(map[string]string)
		students[i]["name"] = fmt.Sprintf("学生%d",i)
		if i%2 == 0 {
			students[i]["sex"] = "男"
		} else {
			students[i]["sex"] = "女"
		}
	}
	fmt.Println(students)

}
