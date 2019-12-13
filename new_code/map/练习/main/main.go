package main

import "fmt"

func main()  {

	/**
	1.使用map[string]map[string]string 的map类型
	2.key表示用户名，是唯一的不可以重复
	3.如果某个用户存在，将其的年龄修改为102，如果不存在创建一个用户年龄是101
	4.通过ModifyUser(user map[string]map[string]string, name string) 去实现
	 */
	var student = make(map[string]map[string]string)

	student["张三"] = make(map[string]string)
	student["张三"]["age"] = "100"
	student["张三"]["sex"] = "男"

	ModifyUser(student,"张三")
	ModifyUser(student,"李四")
	fmt.Println(student)

}

func ModifyUser(user map[string]map[string]string, name string)  {

	if user[name] !=nil {
		user[name]["age"] = "101"
	} else {
		user[name] = make(map[string]string)
		user[name]["age"] = "102"
		user[name]["sex"] = "女"
	}
}