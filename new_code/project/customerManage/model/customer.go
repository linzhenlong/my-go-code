package model

import "fmt"

// 声明一个Customer结构体，表示一个客户信息

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

/**
编写一个工厂模式，返回Customer实例
 */
func NewCustomer(id int,
	name string,
	gender string,
	age int,
	phone string,
	email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomer2(
	name string,
	gender string,
	age int,
	phone string,
	email string) *Customer {
	return &Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
func (customer *Customer)GetInfo() string {
	str := "";
	str += fmt.Sprintf("%v\t", customer.Id)
	str += fmt.Sprintf("%v\t", customer.Name)
	str += fmt.Sprintf("%v\t", customer.Gender)
	str += fmt.Sprintf("%v\t", customer.Age)
	str += fmt.Sprintf("%v\t", customer.Email)
	str += fmt.Sprintf("%v\t", customer.Phone)
	return str
}