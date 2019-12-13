package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var flag bool
	if !flag {
		fmt.Println("is true")
	} else {
		fmt.Println("is false")
	}

	var str string
	fmt.Scanf("%s",&str)

	number,err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number)
	}
}
