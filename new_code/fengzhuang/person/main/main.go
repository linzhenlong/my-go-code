package main

import (
	"fmt"
	personModel "github.com/linzhenlong/my-go-code/new_code/fengzhuang/person/model"
)

func main() {

	p := personModel.NewPerson("TOM")
	fmt.Println(*p)
	p.SetAge(18)
	fmt.Println("年龄是:",p.GetAge())
	p.SetSal(6000.00)
	fmt.Println("薪水是:",p.GetSal())
}