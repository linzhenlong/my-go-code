package main

import (
    "fmt"
    "time"
)

const (
    a = iota
    b
    c
    d
)
func main()  {
    fmt.Println("hello world")
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(time.Now().Unix())
}