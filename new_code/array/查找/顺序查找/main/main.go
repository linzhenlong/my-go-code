package main

import "fmt"

func main()  {

	var names = [4]string{"金毛狮王","紫衫龙王","白眉鹰王","青翼蝠王"}
	var name = ""

	var flag = false
	for true {
		fmt.Println("请输入名称。。。")
		fmt.Scanln(&name)
		for _, val := range names {
			if val == name {
				fmt.Println("success")
				flag = true
				break;
			}
		}
		if flag {
			break
		}

	}

}

func Find()  {

}