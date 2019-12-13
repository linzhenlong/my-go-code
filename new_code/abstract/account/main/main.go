package main

import "fmt"

/**
	定义结构体
 */
type Account struct {
	AccountNo string  //账号
	Pwd string // 密码
	Balance float64 // 余额
}

// 方法
// 1.存款

func (account *Account) SaveMoney(money float64, pwd string) {
	if account.Pwd != pwd {
		fmt.Println("密码错误");
		return
	}
	if money <=0 {
		fmt.Println("存入金额有误");
		return
	}
	account.Balance += money
	fmt.Println("存款成功，当前余额是:",fmt.Sprintf("%.4f",account.Balance))
}

// 取款
func (account *Account) WithDraw(money float64, pwd string) {
	if account.Pwd != pwd {
		fmt.Println("密码错误");
		return
	}
	if money <=0 {
		fmt.Println("存入金额有误");
		return
	}
	if money > account.Balance {
		fmt.Println("取款失败余额不足，当前余额",fmt.Sprintf("%.4f",account.Balance));
		return
	}
	account.Balance -= money
	fmt.Println("取款成功，当前余额是:",fmt.Sprintf("%.4f",account.Balance))
}

func (account *Account)CheckPwd(pwd string) bool {
	return account.Pwd == pwd
}

// 查询余额
func (account *Account)Query(pwd string)  {
	if account.Pwd != pwd {
		fmt.Println("密码错误");
		return
	}
	fmt.Println("当前余额是",fmt.Sprintf("%.4f",account.Balance))
}

func main()  {

	var accountNo1 = Account{
		AccountNo:"No_01",
		Pwd:"123456",
		Balance:19.8,
	}
	i := 1
	for {
		fmt.Print("请输入密码:")
		var pwd string
		_,err :=fmt.Scanf("%s\n",&pwd)
		if err != nil {
			fmt.Println("密码输入有误")
			break
		}
		if accountNo1.CheckPwd(pwd) {
			flag := false
			fmt.Println("选择操作类型:查询(1);取款(2);存款(3)")
			var actionType int
			var money float64
			_,err := fmt.Scanf("%d\n",&actionType)
			if err != nil {
				fmt.Println("操作类型输入有误")
				break
			}
			switch actionType {
			case 1:
				accountNo1.Query(pwd)
			case 2:
				fmt.Println("请输入取款金额...")
				_,err := fmt.Scanf("%f\n",&money)
				if err != nil {
					fmt.Println("取款金额输入有误")
					flag = true
				}
				accountNo1.WithDraw(money,pwd)
			case 3:
				fmt.Println("请输入存款金额...")
				_,err := fmt.Scanf("%f\n",&money)
				if err != nil {
					fmt.Println("存款金额输入有误")
					flag = true
				}
				accountNo1.SaveMoney(money,pwd)
			default:
				fmt.Println("操作类型有误，自动退出")
				flag = true
			}
			if flag {
				break
			}

		} else {
			i++
			if i > 3 {
				fmt.Println("密码输入错误超过三次，请稍后再试")
				break
			}
		}

	}

}
