package main

import (
	"github.com/linzhenlong/my-go-code/day1/packpage_demo/calc"
	"github.com/linzhenlong/my-go-code/oldboy/day8/test/stu"
	"testing"
)

func TestAdd(t *testing.T) {
	// 表格驱动测试
	tests := []struct{ a, b, c int }{
		{1, 2, 3},
		{2, 4, 6},
		{3, 9, 12},
	}
	for _, val := range tests {
		if res := calc.Add(val.a, val.b); res != val.c {
			t.Errorf("(%d + %d) 应该等于%d,但是=%d", val.a, val.b, val.c, res)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{10, 5, 5},
		{8, 3, 5},
		{100, 78, 22},
		{333, 800, -467},
		{16, 8, 8},
	}

	for _, val := range tests {
		if res := calc.Sub(val.a, val.b); res != val.c {
			t.Errorf("(%d - %d) 应该等于%d,但是=%d", val.a, val.b, val.c, res)
		}
	}
}

func TestSave(t *testing.T) {
	student := stu.Student{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}
	err := student.Save()
	if err != nil {
		t.Errorf("stu save err:%s", err.Error())
	}
}

func TestLoad(t *testing.T) {
	student := stu.Student{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}
	res, err := student.Load()
	if err != nil {
		t.Errorf("stu load err:%s", err.Error())
	}
	if student.Name != res.Name {
		t.Errorf("结果不匹配,预期name:%s,实际name:%s", student.Name, res.Name)
	}
	t.Log(res)

}

func TestFlow(t *testing.T) {
	t.Run("save", TestSave)
	t.Run("load", TestLoad)
}

func TestMain(m *testing.M) {
	m.Run()
}
