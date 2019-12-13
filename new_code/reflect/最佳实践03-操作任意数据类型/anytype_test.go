package main

import (
	"reflect"
	"strconv"
	"testing"
)

type Model struct {
	UserId int `json:"user_id"`
	Name string `json:"name"`
}

type Student struct {
	Name string
	Age int
}

func TestAnyType(t *testing.T)  {

	var model *Model

	model = &Model{}
	modelFv := reflect.ValueOf(model)
	t.Log("reflect.ValueOf(model)==>",modelFv)
	t.Log("reflect.ValueOf(model).Kind()的Kind",modelFv.Kind().String())
	// 指针转换
	modelElem := modelFv.Elem()
	t.Log("modelFv.Elem()的Kind",modelElem.Kind())
	modelElem.FieldByName("UserId").SetInt(10)
	t.Log(model)
	modelElem.Field(1).SetString("小明")
	t.Log(model)

	var myMap map[int]Model
	myMap = make(map[int]Model, 10)
	for i:=0;i<10;i++ {
		modelElem.Field(0).SetInt(int64(i))
		modelElem.Field(1).SetString("学生"+strconv.Itoa(i))
		// *model 解引用
		myMap[i] = *model
	}
	t.Log("myMap",myMap)
	myMapFv := reflect.ValueOf(myMap)
	t.Log("myMapFv的Kind()===>",myMapFv.Kind())
	t.Log("myMapFv.MapKeys()==>",myMapFv.MapKeys())
	mayMapKeys := myMapFv.MapKeys()

	t.Log(mayMapKeys[0])

	m2 := &Model{}
	m2Fv := reflect.ValueOf(m2)
	for _,v := range mayMapKeys {

		t.Log(myMapFv.MapIndex(v))
		m2Fv.Elem().FieldByName("Name").SetString("好学生"+strconv.Itoa(int(v.Int())))
		m2Fv.Elem().FieldByName("UserId").SetInt(19+int64(v.Int()))
		myMapFv.SetMapIndex(v,m2Fv.Elem())
	}
	t.Log(myMapFv)

	// 通过反射创建一个结构体
	var stu *Student

	stuFt := reflect.TypeOf(stu)
	// st的指向类型
	st := stuFt.Elem()
	// New 返回一个reflect.Value类型值，该值持有指向类型为type的新申请的0值指针
	elem := reflect.New(st)
	t.Log(elem)
	// stu就是创建的Student的结构体变量
	stu = elem.Interface().(*Student)
	t.Log(stu)
	// 取得elem指向的值
	elem = elem.Elem()
	t.Log(elem)

	// 赋值
	elem.FieldByName("Name").SetString("创建的")
	elem.FieldByName("Age").SetInt(18)
	t.Log(stu)
	t.Log(*stu)

}
