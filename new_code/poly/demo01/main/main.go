package main

import "fmt"

// 声明一个接口
type Usb interface {
	// 声明了两个未实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}
// 让Phone结构体实现Usb接口
func (p Phone)Start()  {
	fmt.Printf("'%s'手机开始工作。。。。\n",p.name)
}

func (p Phone)Stop()  {
	fmt.Printf("'%s'手机停止工作。。。。\n",p.name)
}

func (p Phone)Call()  {
	fmt.Printf("'%s'开始打电话。。。。\n",p.name)
}
// 让Camera实现 Usb接口方法
type Camera struct {
	name string
}
func (c Camera)Start()  {
	fmt.Printf("'%s'相机开始工作。。。。\n",c.name)
}

func (c Camera)Stop()  {
	fmt.Printf("'%s'相机停止工作。。。。\n",c.name)
}

// 计算机
type Computer struct {
	name string
}
// 编写一个working的方法,接收一个usb的接口类型变量
// 只要是实现了Usb接口(就是指实现了Usb接口的所有方法)
func (c *Computer)Working(usb Usb) {
	usb.Start()
	phone,ok := usb.(Phone)
	if ok {
		phone.Call()
	}
	usb.Stop()
}

func main()  {

	// 定义一个Usb接口数组，可以存放phone，camera
	// 这里就体现出多态数组
	var usbArr [3]Usb

	usbArr[0] = Phone{name:"iphone 8"}
	usbArr[1] = Phone{name:"iphone 8s"}
	usbArr[2] = Camera{name:"佳能"}
	fmt.Println(usbArr)


	//先创建结构体变量
	computer := Computer{}
	for _,v := range usbArr {
		computer.Working(v)
	}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}
