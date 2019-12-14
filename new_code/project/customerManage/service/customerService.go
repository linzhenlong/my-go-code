package service

import "github.com/linzhenlong/my-go-code/new_code/project/customerManage/model"

type CustomerService struct {
	customers []model.Customer

	// 新增客户编号 id+1
	CustomerNum int
}

func NewCustomerService() *CustomerService  {

	// 为了测试方便先初始化一个客户
	customerService := &CustomerService{}
	customerService.CustomerNum = 1

	customer := model.NewCustomer(
		1,
		"张三",
		"男",
		18,
		"13800138000",
		"lzl@abc.com")
	customerService.customers = append(customerService.customers, *customer)
	return customerService
}

// 客户列表
func (CustomerService *CustomerService)List()[]model.Customer {
	return CustomerService.customers
}

// 添加客户.
func (CustomerService *CustomerService)Add(name string, gender string, age int,phone string, email string)  {
	id := CustomerService.CustomerNum
	id ++
	customer := model.NewCustomer(id,name,gender,age,phone,email)
	CustomerService.customers = append(CustomerService.customers, *customer)
}

// 添加用户 new
func (CustomerService *CustomerService)AddNew(customer model.Customer) bool {
	CustomerService.CustomerNum++

	customer.Id = CustomerService.CustomerNum
	CustomerService.customers = append(CustomerService.customers, customer)
	return true
}
// 通过客户编号查找客户切片的索引值.
func (CustomerService *CustomerService)FindCustomerIndexById(id int) int {
	index := -1
	for k, v := range CustomerService.customers {
		if v.Id == id {
			index = k
			break
		}
	}
	return index
}

// 删除.
func (CustomerService *CustomerService)Delete(id int) bool {
	index := CustomerService.FindCustomerIndexById(id)
	if index < 0 {
		return  false
	}

	// 从切片中删除一个元素
	CustomerService.customers = append(CustomerService.customers[:index],CustomerService.customers[index+1:]...)
	return true
}

//
func (CustomerService *CustomerService)Edit(id int,customer model.Customer) bool {
	index := CustomerService.FindCustomerIndexById(id)
	if index < 0 {
		return  false
	}

	if customer.Name != "" {
		CustomerService.customers[index].Name = customer.Name
	}
	if customer.Age != 0 {
		CustomerService.customers[index].Age = customer.Age
	}
	if customer.Gender != "" {
		CustomerService.customers[index].Gender = customer.Gender
	}
	if customer.Phone != "" {
		CustomerService.customers[index].Gender = customer.Phone
	}
	if customer.Email != "" {
		CustomerService.customers[index].Gender = customer.Email
	}


	return true
}