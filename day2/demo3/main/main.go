package main

import (
    "fmt"
    "go_dev/day2/demo3/sex"
    "time"
)

func main() {
    for {
        var s string = sex.Print()
        fmt.Println(s)
        time.Sleep(1000 * time.Millisecond)
    }
}