package main

import (
    "fmt"
    "github.com/linzhenlong/my-go-code/day1/packpage_demo/calc"
)
func main()  {
    sum := calc.Add(100,200)
    fmt.Println("sum=",sum)
    sub := calc.Sub(100,300)
    fmt.Println("sub=",sub)
}
