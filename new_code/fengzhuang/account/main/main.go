package main

import (
	"fmt"
	"go_dev/new_code/fengzhuang/account/model"
)

func main()  {

	tag := make(map[string]string)
	tag["1"] = "骑马"
	tag["2"] = "射箭"

	account := model.NewAccount("100000","tom",tag,2000.00,"123456")
	if account == nil {
		return
	}
	fmt.Println("skill=",account.GefSkill())
	tag["3"] = "踢足球"
	account.SetSkill(tag)
	fmt.Println("skill=",account.GefSkill())
	fmt.Println("原始密码是:",account.GetPwd())
	account.SetPwd("56789")
	fmt.Println("修改后的密码是:",account.GetPwd())
}