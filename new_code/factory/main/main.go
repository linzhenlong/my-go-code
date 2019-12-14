package main

import (
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/factory/model"
)

func main()  {

	/*var stu model.Student
	stu.Name = "TOM"
	fmt.Println(stu.Name)
	stu.Say()*/

	//当student结构体首字母小写,我们可以通过工厂模式解决
	var stu = model.NewStudent("TOM",98.8)
	fmt.Println(*stu)
	stu.Say()
	// 因为字段score首字母小写，则在其他包不可以直接使用
	fmt.Println(stu.GetScore())
}
