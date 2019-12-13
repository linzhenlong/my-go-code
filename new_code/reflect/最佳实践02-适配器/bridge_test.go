package main

import (
	"reflect"
	"testing"
)

func TestBridge(t *testing.T) {
	num := Bridge(Test1,1,9)
	res ,ok:= num.(int)
	if !ok {
		t.Fatalf("类型有误,正确的类型是%v,现在是:%T\n",reflect.ValueOf(num).Kind(),res)
	}
	if res != 10 {
		t.Fatalf("计算错误，预期值10，实际值%v\n",res)
	}
	t.Log("Bridge(Test1,1,9)测试通过")

	num2 := Bridge(Test2,2,2,"*")
	num2 = num2.(float64)
	if Test2(2,2,"*") != num2 {
		t.Fatalf("Bridge(Test2,2,2,\"*\"),与实际值不符\n")
	}

	t.Logf("Bridge(Test2,2,2,\"*\")=%v\n",num2)

}
