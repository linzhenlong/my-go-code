package main

import "fmt"

// 声明一个接口
type Usb interface {
	// 声明了两个未实现的方法
	Start()
	Stop()
}

type Phone struct {

}
// 让Phone结构体实现Usb接口
func (p Phone)Start()  {
	fmt.Println("手机开始工作。。。。")
}

func (p Phone)Stop()  {
	fmt.Println("手机停止工作。。。。")
}
// 让Camera实现 Usb接口方法
type Camera struct {

}
func (c Camera)Start()  {
	fmt.Println("相机开始工作。。。。")
}

func (c Camera)Stop()  {
	fmt.Println("相机停止工作。。。。")
}

// 计算机 
type Computer struct {

}
// 编写一个working的方法,接收一个usb的接口类型变量
// 只要是实现了Usb接口(就是指实现了Usb接口的所有方法)
func (c *Computer)Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main()  {
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}
