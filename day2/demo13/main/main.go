package main

import (
    "math/rand"
    "fmt"
    "time"
)

func init() {
    // 该函数设置随机种子
    // 若不调用此函数设置随机种子，则默认的种子值为1，由于随机算法是固定的，
    // 如果每次都以1作为随机种子开始产生随机数，则结果都是一样的，因此一般
    // 都需要调用此函数来设置随机种子，通常的做法是以当前时间作为随机种子
    // 以保证每次随机种子都不同，从而产生的随机数也不通
    // 该函数协程安全
    rand.Seed(time.Now().Unix())
}
func main()  {
    fmt.Println("生成10个小于100的随机数")
    for i:=0;i<10;i++ {
        a := rand.Int31n(100)
        fmt.Println(a)
    }
    fmt.Println("生成10个随机的浮点数")
    for i:=0;i<10;i++ {
        a :=rand.Float32()
        fmt.Println(a)
    }
}
