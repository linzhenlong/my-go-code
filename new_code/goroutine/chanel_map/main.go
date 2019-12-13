package main

import "fmt"

// 管道里存放map

func main()  {
	chanMap := make(chan map[string]string, 3)

	myMap1 := make(map[string]string , 3)

	myMap1["name"] = "张三"
	myMap1["age"] = "18"

	chanMap <- myMap1

	myMap2 := make(map[string]string , 3)
	myMap2["name"] = "李四"
	myMap2["age"] = "20随"

	chanMap <- myMap2

	fmt.Println(<-chanMap)
	fmt.Println(<-chanMap)
}
