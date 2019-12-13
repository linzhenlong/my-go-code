package main

import "fmt"

func main() {
    var a int = 100
    var b chan int = make(chan int,100)
    fmt.Println("a=",a)  // 100
    modify(a)
    fmt.Println("a=",a) // 100
    modif1(&a)
    fmt.Println("a=",a)  // 10
    fmt.Println("b=",b) // 0xc42009c000
}

func modify(a int) {
    a = 10
    return
}

func modif1(a *int) {
    *a = 10
}
