package model

import "fmt"

//因为student 结构体首字母是小写，因此相当于是private,只能本包内使用
//我们通过工厂模式来解决这个问题
type student struct {
	Name string
	score float64  // 如果字段score首字母小写，则在其他包不可以直接使用，可以通过func (stu *student) GetScore () float64{...}实现
}

func NewStudent(Name string,Score float64) *student{
	return &student{
		Name:Name,
		score:Score,
	}
}
func (stu *student) Say () {
	fmt.Println("I AM ",stu.Name)
}

func (stu *student) GetScore () float64{
	return stu.score
}