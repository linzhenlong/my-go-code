package main

import (
	monster2 "go_dev/new_code/testing/testcase02/monster"
	"testing"
)



func TestReStore(t *testing.T)  {
	// 先创建一个monster 实例，先不需要指定字段值
	var monster = monster2.NewMonster("", 0,[]string{})

	res , err := monster.ReStore()
	if err !=nil {
		t.Fatalf("TestReStore 错误%v\n",err)
	}
	if !res {
		t.Fatalf("TestReStore 读取错误\n")
	}
	if monster.Name != "牛魔王" {
		t.Fatalf("TestReStore 读取错误,(*monster).Name实际值%s,期望值%s\n",monster.Name,"牛魔王")
	}
	t.Logf("TestReStore success name:%s;age:%d,skill:%v\n", monster.Name,monster.Age,monster.Skill)
}
func TestStore(t *testing.T)  {
	var skill = []string{"1111","2222"}
	monster :=  monster2.NewMonster("牛魔王",100,skill)
	f , err := monster.Store()
	if err !=nil {
		t.Fatalf("错误%v\n",err)
	}
	if !f {
		t.Fatalf("写入错误错误\n")
	}
	t.Logf("success\n")
}