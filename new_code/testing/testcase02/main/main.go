package main

import (
	"fmt"
	m "github.com/linzhenlong/my-go-code/new_code/testing/testcase02/monster"
)

func main()  {

	var skill = []string{"1111","2222"}
	 monster :=  m.NewMonster("牛魔王",100,skill)


	f , err := (*monster).Store()
	fmt.Printf("%v %v", f,err)


}
