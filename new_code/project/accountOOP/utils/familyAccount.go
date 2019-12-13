package utils

import (
	"fmt"
	"strings"
)

// 定义结构体
type familyAccount struct {
	// 声明变量保存用户选择
	key string

	// 声明变量是否退出
	isExit bool

	/*// 定义账号余额
	var balance float64

	// 每次收支的金额
	money := 0.0*/

	// 每次收支的说明
	balance float64 // 定义账号余额
	money   float64 // 每次收支的金额
	note    string  // 收支说明
	details string  // 收支详情

	//details = "收支\t账户金额\t收支金额\t说明\n"

	// 定义一个变量标识是否有收入或是支出
	isHasInOut bool
}

// 构造方法
func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		balance: 1000.00,
		isExit:false,
		money:0.0,
		isHasInOut:false,
		details:"收支\t账户金额\t收支金额\t说明\n",
	}
}


// 显示收支明显
func (account *familyAccount) showDetails() {
	if account.isHasInOut {
		title := "----------收支明细----------------\n"
		title += account.details
		fmt.Println(title)
	} else {
		fmt.Println("当前还没有收入或是支出，请来一笔吧...")
	}
}

// 显示支出
func (account *familyAccount) out() {
START:
	fmt.Print("请输入支出金额:")
	_, _ = fmt.Scanln(&account.money)
	if (account.balance - account.money) < 0 {
		fmt.Print("输入支出金额有误，")
		goto START
	}
	fmt.Print("请输入支出说明:")
	_, _ = fmt.Scanln(&account.note)
	account.balance -= account.money
	account.details += fmt.Sprintf("支出\t%.2f\t%.2f\t%s\n", account.balance, account.money, account.note)
	fmt.Println("----------登记完成---------------")
	account.isHasInOut = true
}

// 将登记收入写成方法.
func (account *familyAccount) inCome() {
	fmt.Print("请输入收入金额:")
	_, _ = fmt.Scanln(&account.money)
	fmt.Print("请输入收入说明:")
	_, _ = fmt.Scanln(&account.note)
	account.balance += account.money
	account.details += fmt.Sprintf("收入\t%.2f\t%.2f\t%s\n", account.balance, account.money, account.note)
	fmt.Println("----------登记完成---------------")
	account.isHasInOut = true
}

// 退出系统
func (account *familyAccount) exitSystem() {
	fmt.Print("您确定要退出吗？(y/n)")
	choose := ""
	for {
		_, _ = fmt.Scanln(&choose)
		if strings.ToLower(choose) == "y" || strings.ToLower(choose) == "n" {
			break
		} else {
			fmt.Print("退出请输入，(y/n)")
		}
	}
	if strings.ToLower(choose) == "y" {
		account.isExit = true
	}
}

// 显示主菜单方法
func (account *familyAccount) MainMenu() {
	for {
		str := "----------家庭收支明细---------------\n"
		str += "		    1.收支明细\n"
		str += "		    2.登记收入\n"
		str += "		    3.登记支出\n"
		str += "		    4.退   出 \n"
		str += "		    请选择（1-4）:"
		fmt.Print(str)
		_, _ = fmt.Scanln(&account.key)
		switch account.key {
		case "1":
			account.showDetails()
		case "2":
			account.inCome()
		case "3":
			account.out()
		case "4":
			account.exitSystem()

		default:
			fmt.Println("请输入正确的选项")
		}
		if account.isExit {
			break
		}
	}
}
