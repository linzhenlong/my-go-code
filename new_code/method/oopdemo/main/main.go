package main

import "fmt"

type Student struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Id int `json:"id"`
	Score float64 `json:"score"`
}

func (student *Student) Say () {
	str := fmt.Sprintf("我是%s,我是%s的,我的学号是%d,我今年%d岁了,期末考试考了%.2f分\n",
		(*student).Name,student.Gender,student.Id,student.Age,student.Score)
	fmt.Println(str)
}

type Box struct {
	Length float64
	Width float64
	Height float64
}

func (box *Box) volume() float64 {
	return box.Height * box.Width * box.Length
}

type Visitor struct {
	Name string
	Age int
}

func (visitor *Visitor) ShowPrice()  {
	if visitor.Age > 18 {
		fmt.Println("门票价格",20)
	} else {
		fmt.Println("门票价格",10)
	}
}
func main()  {

	student1 := Student{
		Name:"TOM",
		Gender:"女",
		Age:18,
		Id:100,
		Score:95.499,
	}

	student1.Say()

	var box Box
	fmt.Println("请输入长...")
	fmt.Scanf("%f\n",&(box.Length))

	fmt.Println("请输入宽...")
	fmt.Scanf("%f\n",&(box.Width))

	fmt.Println("请输入高...")
	fmt.Scanf("%f\n",&(box.Height))

	fmt.Println("面积是",box.volume())

	for {
		var name string
		var age int
		fmt.Print("姓名:")
		fmt.Scanf("%s\n",&name)
		if name == "n" {
			fmt.Println("退出了....")
			break
		}
		fmt.Print("年龄:")
		fmt.Scanf("%d\n",&age)

		var visitor Visitor
		visitor.Age = age
		visitor.Name = name
		visitor.ShowPrice()

		fmt.Println("输入名字n\n退出程序。。。")
	}

}
