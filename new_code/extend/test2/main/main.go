package main

import "fmt"

// 学生
type Student struct {
	Name string
	Age int
	Score float64
}

// 将pupil,graduate共有的ShowInfo方法绑定 给*Student
func (stu *Student)ShowInfo() {
	fmt.Printf("学生姓名:%s,年龄:%d岁,得分:%.2f \n",stu.Name,stu.Age,stu.Score)
}

// 将pupil,graduate共有的SetScore方法绑定 给*Student
func (stu *Student) SetScore(score float64) {
	stu.Score = score
}
// 将pupil,graduate共有的GetScore方法绑定 给*Student
func (stu *Student)GetScore() float64 {
	return stu.Score
}
// 小学生
type Pupil struct {
	Student // 嵌套匿名结构体,实现基础student
	FlowerNum int // 小红花个数,小学生独有
}

// 小学生特有方法
func (pupil *Pupil)TestIng() {
	fmt.Println("小学生",pupil.Name,"正在考试")
}

// 大学生特有方法
func (pupil *Graduate)TestIng() {
	fmt.Println("大学生",pupil.Name,"正在考试")
}
// 大学生
type Graduate struct {
	Student // 嵌套匿名结构体,实现基础student
	Bursary float64 // 奖学金
}

func main()  {

	// 当我们对结构体迁入匿名结构体后，使用方法:

	pupil := &Pupil{}

	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.Student.Score = 97.00
	pupil.FlowerNum = 10
	pupil.TestIng()
	pupil.SetScore(60)
	fmt.Printf("小学生%s,考了%.2f 分\n",pupil.Student.Name,pupil.Student.GetScore())
	pupil.ShowInfo()


	var graduate Graduate
	graduate.Name = "大学生" // 赋值是省略.Student也可以,充分体现继承
	graduate.Age = 20
	graduate.Score = 100
	graduate.ShowInfo()
	graduate.TestIng()
	graduate.SetScore(36)
	fmt.Printf("大学生%s,考了%.2f 分\n",graduate.Name,graduate.GetScore())
}
