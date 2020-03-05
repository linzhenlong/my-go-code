package main

import "fmt"

func main()  {
    var str string = "hello world\n"
    var str2 string = `hello world \n`

    var b byte = 'c'
    fmt.Println("str=",str)
    fmt.Println("str2=",str2)
    fmt.Println(b)
    fmt.Printf("%c\n",b)
}
