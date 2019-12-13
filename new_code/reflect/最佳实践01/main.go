package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 定义了一个person结构体
type Person struct {
	Name string `json:"persoon_name"`
	Age int `json:"age"`
	Sex string `json:"sex"`
	Phone string `json:"dianhua"`
}

func (person *Person)Print()  {
	fmt.Println("Print start...")
	fmt.Println(person)
	fmt.Println("Print end...")
}

func (person *Person)GetSum(n1 int, n2 int) int  {
	return n1+n2
}

func (person *Person)Set(name string,age int,sex string, phone string )  {
	person.Name = name
	person.Age = age
	person.Sex = sex
	person.Phone = phone
}


func TestStruct(val interface{})  {
	fv := reflect.ValueOf(val)
	ty := reflect.TypeOf(val)
	fmt.Printf("fv.Kind()=%v\n",fv.Kind())
	fmt.Printf("fv.Elem().Kind()=%v\n",fv.Elem().Kind())

	if fv.Elem().Kind() != reflect.Struct {
		fmt.Println("类别有误,kind为",fv.Kind())
		return
	}
	fmt.Println(fv.Elem().Kind())
	// 这块不需要用Elem()应该是因为方法是结构体指针绑定的
	methodNums := fv.NumMethod()
	fmt.Printf("Person结构体绑定了%d个方法\n",methodNums)

	// 获取结构体的字段个数，由于传入的是指针类型，因此需要Elem()转换一下.
	numFields := fv.Elem().NumField()
	fmt.Printf("Person结构体有%d个字段\n",numFields)
	for i:=0;i<numFields;i++ {
		fmt.Printf("Person结构体有 field %d的值为:%v\n",i,fv.Elem().Field(i))
		// 获取结构体字段标签的值,获取标签需要用reflect.Type来获取Tag值
		tagVal := ty.Elem().Field(i).Tag.Get("json")

		if tagVal !="" {
			fmt.Printf("Person结构体有 field %d标签值为:%v\n",i,tagVal)
		}
	}

	// 调用结构体方法,如果方法木有参数，那么Call()传参nil
	// fv.Method(1).Call(nil)调用的第二个方法，但是这里的排序并不是按照，代码里的方法顺序排序的
	// 是按照方法名称的ASCII码进行排序的，所以Method(1)对应的是Print()的方法
	fv.Method(1).Call(nil)

	// Call()有参数的方法时，Call参数的的类型为[]reflect.Value
	var params1 []reflect.Value
	params1 = append(params1,reflect.ValueOf(10))
	params1 = append(params1,reflect.ValueOf(20))

	resGetSum := fv.MethodByName("GetSum").Call(params1)
	fmt.Printf("%v\n",reflect.ValueOf(resGetSum))
	for _,v := range resGetSum {
		fmt.Printf("(person *Person)GetSum(n1 int, n2 int)的结果为:%d\n",v)
	}

	// 调用Set()方法
	var params2 []reflect.Value
	params2 = append(params2,reflect.ValueOf("jack"))
	params2 = append(params2,reflect.ValueOf(22))
	params2 = append(params2,reflect.ValueOf("女"))
	params2 = append(params2,reflect.ValueOf("1755555"))
	fv.Method(2).Call(params2)
	fmt.Printf("%v\n",fv)

	// 在修改一下person 的字段值
	fv.Elem().Field(0).SetString("jack2")
	fv.Elem().FieldByName("Age").SetInt(19)

}
func main()  {

	person := &Person{
		Name:"tom",
		Age:18,
		Sex:"男",
		Phone:"13800138000",
	}
	TestStruct(person)
	fmt.Printf("main中person:%v\n",*person)

	jsonBytes,err := json.Marshal(*person)
	if err != nil {
		return
	}
	fmt.Println(string(jsonBytes))
}
