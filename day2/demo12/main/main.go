package main

import "fmt"

func main()  {
    var n int16 = 34
    var m int32
    //m = n // 报错cannot use n (type int16) as type int32 in assignment
    m = int32(n)
    fmt.Printf("32bit int is %d\n",m)
    fmt.Printf("16bit int is %d\n",n)

}
