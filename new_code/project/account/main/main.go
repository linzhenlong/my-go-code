package main

import (
	"fmt"
	"strings"
)

func main() {

	// 声明变量保存用户选择
	key := ""

	// 声明变量是否退出
	var isExit bool
	isExit = false

	/*// 定义账号余额
	var balance float64

	// 每次收支的金额
	money := 0.0*/

	// 每次收支的说明
	var (
		balance float64 // 定义账号余额
		money   float64 // 每次收支的金额
		note    string  // 收支说明
		details string  // 收支详情
	)
	balance = 10000.0 // 默认账户里有10000元

	details = "收支\t账户金额\t收支金额\t说明\n"

	// 定义一个变量标识是否有收入或是支出
	isHasInOut := false
	// 显示主菜单
	for {
		mainFace()
		_,_ = fmt.Scanln(&key)
		switch key {
		case "1":
			if isHasInOut {
				incomeDetail(details)
			} else {
				fmt.Println("当前还没有收入或是支出，请来一笔吧...")
			}
		case "2":
			fmt.Print("请输入收入金额:")
			_, _ = fmt.Scanln(&money)
			fmt.Print("请输入收入说明:")
			_, _ = fmt.Scanln(&note)
			balance += money
			details += fmt.Sprintf("收入\t%.2f\t%.2f\t%s\n", balance, money, note)
			fmt.Println("----------登记完成---------------")
			isHasInOut = true
		case "3":
		START:
			fmt.Print("请输入支出金额:")
			_, _ = fmt.Scanln(&money)
			if (balance - money) < 0 {
				fmt.Print("输入支出金额有误，")
				goto START
			}
			fmt.Print("请输入支出说明:")
			_, _ = fmt.Scanln(&note)
			balance -= money
			details += fmt.Sprintf("支出\t%.2f\t%.2f\t%s\n", balance, money, note)
			fmt.Println("----------登记完成---------------")
			isHasInOut = true
		case "4":

			fmt.Print("您确定要退出吗？(y/n)")
			choose := ""
			for {
				_,_ = fmt.Scanln(&choose)
				if strings.ToLower(choose) == "y" || strings.ToLower(choose) == "n" {
					break
				} else {
					fmt.Print("退出请输入，(y/n)")
				}
			}
			if strings.ToLower(choose) == "y" {
				isExit = true
			}

		default:
			fmt.Println("请输入正确的选项")
		}
		if isExit {
			break
		}
	}
	fmt.Println("已经成功退出....")
}

func mainFace() {
	str := "----------家庭收支明细---------------\n"
	str += "		    1.收支明细\n"
	str += "		    2.登记收入\n"
	str += "		    3.登记支出\n"
	str += "		    4.退   出 \n"
	str += "		    请选择（1-4）:"
	fmt.Print(str)
}

func incomeDetail(detail string) {
	title := "----------收支明细----------------\n"
	title += detail
	fmt.Println(title)
}

