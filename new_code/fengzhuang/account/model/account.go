package model

import "fmt"

type account struct {
	id string
	name string
	skill map[string]string
	pwd string
	money float64
}

func (account *account)SetPwd(pwd string)  {
	account.pwd = pwd
}
func (account *account)GetPwd() string  {
	return account.pwd
}

func (account *account)SetSkill(skill map[string]string)  {
	account.skill = skill
}
func (account *account)GefSkill() map[string]string {
	return account.skill
}

func NewAccount(id string,name string,skill map[string]string,money float64,pwd string) *account {
	if len(id) < 6 || len(id) > 15 {
		fmt.Println("id 长度有误")
		return nil
	}
	if name == "" {
		fmt.Println("name 不能为空")
	}
	return &account{
		name:name,
		pwd:pwd,
		skill:skill,
		money:money,
	}
}
