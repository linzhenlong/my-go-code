package main

import "fmt"

func main() {
    var first int = 100
    var second int = 200
    second, first = first, second
    /*
        first,second = swap(first,second)
    */

    fmt.Println("first=",first)
    fmt.Println("second=",second)
}

func swap(a int, b int) (int, int) {
    return b, a
}