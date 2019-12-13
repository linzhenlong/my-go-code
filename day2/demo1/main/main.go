package main

import "fmt"

func main()  {
    add(5)
}
func add(n int) {
    for i:=0;i<=n;i++ {
        fmt.Printf("%d+%d=%d\n",i,(n-i),n)
    }
}