package main

import (
	"fmt"
)

type student struct {
	name string
}

func main() {
	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)

	stu := student{name: "张三"}
	stuChan <- &stu
	var stu1 interface{}
	stu1 = <-stuChan

	var stu02 *student
	stu02, ok := stu1.(*student)
	if ok {
		fmt.Println(stu02)
	}

}
