package view

import (
	"fmt"
	"github.com/linzhenlong/my-go-code/new_code/project/customerManage/model"
	"github.com/linzhenlong/my-go-code/new_code/project/customerManage/service"
	"strings"
)

type customerView struct {
	Key string
	IsExist bool

	//
	CustomerService *service.CustomerService
	name string
	gender string
	age int
	email string
	phone string
}

func NewCustomerView() *customerView  {
	return &customerView{
		Key:"",
		IsExist:false,
		CustomerService:service.NewCustomerService(),
	}
}

func (c *customerView)list()  {
	list := c.CustomerService.List()
	fmt.Println("--------------------客户列表-------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t邮箱\t\t\t电话\t\t\t")
	for key,_ := range list {
		fmt.Println(list[key].GetInfo())
	}
	fmt.Println("--------------------客户列表完成----------------")
}

func (c *customerView)add()  {

	fmt.Println("					  添加新客户                        ")
	fmt.Print("名称:")
	_, _ = fmt.Scanln(&c.name)

	fmt.Print("性别:")
	_, _ = fmt.Scanln(&c.gender)


	fmt.Print("年龄:")
	_, _ = fmt.Scanln(&c.age)

	fmt.Print("电话:")
	_, _ = fmt.Scanln(&c.phone)

	fmt.Print("email:")
	_, _ = fmt.Scanln(&c.email)


	customer := model.NewCustomer2(c.name,c.gender,c.age, c.phone,c.email)
	f := c.CustomerService.AddNew(*customer)
	if f {
		fmt.Println("					  添加完成                        ")
	} else {
		fmt.Println("					  添加失败                        ")
	}
	/*c.CustomerService.Add(
		c.name,
		c.gender,
		c.age,
		c.phone,
		c.email,
		)*/
}

func (c *customerView)delete()  {
	fmt.Println("---------------------删除客户-----------------")
	fmt.Print("请输入待删除客户的编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := c.CustomerService.FindCustomerIndexById(id)
	if index == -1 {
		fmt.Println("用户不存在")
		return
	}
	customer := c.CustomerService.List()
	customerInfo := customer[index].GetInfo()
	fmt.Print("确认删除客户："+customerInfo+"吗？（Y/N）")
	choose := ""
	for {
		fmt.Scanln(&choose)
		if strings.ToLower(choose) =="y" || strings.ToLower(choose) =="n" {
			break
		} else {
			fmt.Println("确认删除"+customerInfo+"吗？（Y/N）")
		}
	}

	if strings.ToLower(choose) =="y" {
		if c.CustomerService.Delete(id) {
			fmt.Println("删除成功")
		} else {
			fmt.Println("删除失败")
		}
	} else {
		return
	}
}

func (c *customerView)edit()  {
	fmt.Println("---------------------编辑客户-----------------")
	fmt.Print("请输入待编辑客户的编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := c.CustomerService.FindCustomerIndexById(id)
	if index == -1 {
		fmt.Println("用户不存在")
		return
	}
	customer := c.CustomerService.List()
	customerInfo := customer[index].GetInfo()
	fmt.Print("确认编辑客户："+customerInfo+"吗？（Y/N）")
	choose := ""
	for {
		fmt.Scanln(&choose)
		if strings.ToLower(choose) =="y" || strings.ToLower(choose) =="n" {
			break
		} else {
			fmt.Println("确认编辑"+customerInfo+"吗？（Y/N）")
		}
	}
	if strings.ToLower(choose) =="y" {
		fmt.Print("修改名称（"+customer[index].Name+"）:")
		_, _ = fmt.Scanln(&c.name)

		fmt.Print("修改性别（"+customer[index].Gender+"）:")
		_, _ = fmt.Scanln(&c.gender)


		fmt.Print("修改年龄（"+fmt.Sprintf("%v",customer[index].Age)+"）:")
		_, _ = fmt.Scanln(&c.age)

		fmt.Print("修改电话（"+customer[index].Phone+"）:")
		_, _ = fmt.Scanln(&c.phone)

		fmt.Print("修改邮箱（"+customer[index].Email+"）:")
		_, _ = fmt.Scanln(&c.email)


		customer := model.NewCustomer2(c.name,c.gender,c.age, c.phone,c.email)
		f := c.CustomerService.Edit(id,*customer)
		if f {
			fmt.Println("					  编辑完成                        ")
		} else {
			fmt.Println("					  编辑失败                        ")
		}
	} else {
		return
	}


}

//显示主菜单

func (c *customerView) ShowMainMenu() {
	for {
		fmt.Println("--------------------客户信息管理系统-------------------")
		fmt.Println()
		fmt.Println("					  1 添加客户                        ")
		fmt.Println("					  2 修改客户						   ")
		fmt.Println("					  3 删除客户						    ")
		fmt.Println("					  4 客户列表						    ")
		fmt.Println("					  5 退出                             ")
		fmt.Println()
		fmt.Println("--------------------请选择（1-5）----------------------")
		_, _ = fmt.Scanln(&c.Key)
		switch c.Key {
		case "1":
			c.add()
		case "2":
			c.edit()
		case "3":
			c.delete()
		case "4":
			c.list()
		case "5":
			c.IsExist = true
		default:
			println("--------------------输入有误，请选择（1-5）----------------------")

		}
		if c.IsExist {
			break
		}
	}
}