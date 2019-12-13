package main

import "testing"

func TestSum2(t *testing.T)  {
	res := sum(1,2)
	if res != 3 {
		t.Fatalf("错误")
	}
	t.Logf("正确")
}
