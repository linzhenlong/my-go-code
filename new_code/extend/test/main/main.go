package main

import "fmt"

// 编写学生考试系统
// 小学生要考试

type Pupil struct {
	Name string
	Age int
	Score float64
}


func (pupil *Pupil)ShowInfo() {
	fmt.Printf("学生姓名:%s,年龄:%d岁,得分:%.2f \n",pupil.Name,pupil.Age,pupil.Score)
}

func (pupil *Pupil)SetScore(score float64) {
	pupil.Score = score
}
func (pupil *Pupil)GetScore() float64 {
	return pupil.Score
}

func (pupil *Pupil)TestIng() {
	fmt.Println("小学生",pupil.Name,"正在考试")
}


// 大学生相关的


type Graduate struct {
	Name string
	Age int
	Score float64
}

func (pupil *Graduate)ShowInfo() {
	fmt.Printf("学生姓名:%s,年龄:%d岁,得分:%.2f \n",pupil.Name,pupil.Age,pupil.Score)
}

func (pupil *Graduate)SetScore(score float64) {
	pupil.Score = score
}
func (pupil *Graduate)GetScore() float64 {
	return pupil.Score
}

func (pupil *Graduate)TestIng() {
	fmt.Println("大学生",pupil.Name,"正在考试")
}

/**
	问题:
		1.小学生跟大学生结构体字段基本一致
		2.方法一致
		3.代码冗余严重
 */


func main()  {
	pupil := Pupil{
		Name:"小明",
		Age:9,
	}
	pupil.ShowInfo()
	pupil.TestIng()
	pupil.SetScore(98.50)
	fmt.Println(pupil.Name,"考了",pupil.GetScore())

	fmt.Println("==================")
	graduate := Graduate{
		Name:"mary",
		Age:9,
	}

	graduate.ShowInfo()
	graduate.TestIng()
	graduate.SetScore(98.50)
	fmt.Println(graduate.Name,"考了",graduate.GetScore())
}
