package main

import (
    "go_dev/day1/goroute_demo/goroute"
    "fmt"
)

func main()  {
    var pipe chan int
    pipe = make(chan int ,1)
    go goroute.Add(100,200,pipe)

    sum := <- pipe
    fmt.Println("sum=",sum)
}
