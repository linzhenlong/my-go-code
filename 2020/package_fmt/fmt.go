package main

import "fmt"

import "os"

import "math"

type User struct {
	name string
	age int
}

func main() {
	fmt.Print("在终端打印信息，hello world......\n")
	name := "沙河小王子"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示....")

	fmt.Println("Fprint.........")
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")

	file, err := os.OpenFile("./abc.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr,"os.OpenFile error :=%s\n", err.Error())
	}
	defer file.Close()
	fmt.Fprintf(file, "往文件中写信息:%s\n", name)

	fmt.Println("Sprint......")
	s1 := fmt.Sprintf("沙河小王子")
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)

	fmt.Println("格式化占位符.......")
	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", false)
	user := User{
		name: name,
		age: age,
	}
	fmt.Printf("%v\n", user)
	fmt.Printf("%#v\n",user)
	fmt.Printf("%T\n",user)
	fmt.Printf("100%%\n")

	fmt.Println("整数占位符.......")
	n :=65
	fmt.Printf("%d转二进制:%b\n",n,n)
	fmt.Printf("%d转unicode编码值:%c\n",n,n)
	fmt.Printf("%d转十进制:%d\n",n,n)
	fmt.Printf("%d转八进制:%o\n",n,n)
	fmt.Printf("%d转八进制:%o\n",n,n)
	fmt.Printf("%d转十六进制，a-f:%x\n",n,n)
	fmt.Printf("%d转十六进制，A-F:%X\n",n,n)

	fmt.Println("浮点数与复数占位符.....")
	fmt.Printf("%b\n",math.Pi)
	fmt.Printf("%e\n", math.Pi)
	fmt.Printf("%E\n", math.Pi)
	fmt.Printf("%f\n",math.Pi)
	fmt.Printf("%.3f\n", math.Pi)
	fmt.Printf("%g\n", math.Pi)
	fmt.Printf("%G\n", math.Pi)

	fmt.Println("字符串和[]byte 占位符....")
	
	s := "古丽小娜扎"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)

	fmt.Println("指针的标识符.....")
	a := 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)
	
	fmt.Println("宽度标识符....")
	fmt.Printf("%f\n", math.Pi) // 默认宽度，默认精度
	fmt.Printf("%9f\n", math.Pi) // 宽度9，默认精度
	fmt.Printf("%.2f\n", math.Pi) // 默认宽度,精度2
	fmt.Printf("%9.2f\n", math.Pi) // 宽度9,精度2
	fmt.Printf("%9.f\n", math.Pi) // 宽度9,精度0

	fmt.Println("其他标识符....")
	s2 = "其他占位符"
	age = -100
	fmt.Printf("%s\n", s2)
	fmt.Printf("%+d\n", age)
	fmt.Printf("%5s\n", s2)
	fmt.Printf("%5s\n", s2)
	fmt.Printf("%-5s\n", s2)
	fmt.Printf("%5.7s\n", s2)
	fmt.Printf("%-5.7s\n", s2)
	fmt.Printf("%5.2s\n", s2)
	fmt.Printf("%05s\n", s2)
}