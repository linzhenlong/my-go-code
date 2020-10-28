package main

import "testing"

func TestAdd(t *testing.T) {
	a, b := 19,20
	res := Add(a,b)
	if res != 40 {
		t.Fatalf("Add(%d,%d)=%d与预期(%d)不符",a,b,res,40)
	} else {
		t.Logf("Add(%d,%d)=%d与预期(%d)不符:",a,b,res,40)
	}
}

func TestJian(t *testing.T) {
	a, b := 50,30
	res := Jian(a,b)
	if res != 20 {
		t.Fatalf("Jian(%d,%d)=%d与预期(%d)不符",a,b,res,20)
	} else {
		t.Logf("Jian(%d,%d)=%d与预期(%d)不符:",a,b,res,20)
	}
}