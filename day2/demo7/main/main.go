package main

import (
    "fmt"
    "os"
)
func main()  {
    var goos string = os.Getenv("")
    fmt.Printf("当前操作系统%s\n",goos)
    path := os.Getenv("GOROOT")
    fmt.Printf("环境变量%s",path)
}
