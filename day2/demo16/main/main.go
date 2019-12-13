package main

import "fmt"

// https://go-zh.org/pkg/fmt/
func main()  {
    var (
        a int
        b bool
        c byte = 'a'
    )
    fmt.Println(`%v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名`)
    fmt.Printf("%v\n",a)
    fmt.Printf("%v\n",b)
    fmt.Printf("%v\n",c)
    fmt.Println(`%#v	相应值的Go语法表示`)
    fmt.Printf("%#v\n",a)
    fmt.Printf("%#v\n",b)
    fmt.Printf("%#v\n",c)
    fmt.Println(`%T	相应值的类型的Go语法表示`)
    fmt.Printf("%T\n",a)
    fmt.Printf("%T\n",b)
    fmt.Printf("%T\n",c)

}