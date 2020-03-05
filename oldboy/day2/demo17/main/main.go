package main

import (
    "fmt"
    _"strings"
)

func main()  {
    var(
        str1 string = "hello"
        str2 string = "world";
    )
    str3 := str1 + " "+ str2
    fmt.Println(str3)

    str4 := fmt.Sprintf("%s %s \n",str1, str2)
    fmt.Println(str4)

    n := len(str3)
    fmt.Printf("len(str3)=%d\n",n)

    // 切片
    substr := str3[0:5]
    fmt.Println(substr)
    reverse_str := reverse(str3)
    fmt.Println(reverse_str)

    fmt.Println(revrese1(reverse_str))
}
/**
   字符串反转
 */
func reverse(str string) string {
    var res string
    len := len(str)
    for i:=0;i<len;i++ {
        res = res + fmt.Sprintf("%c",str[len-i-1])
    }
    return res
}
func revrese1(str string) string {
    var res []byte
    tmp := []byte(str)
    len := len(str)
    for i:=0;i<len;i++ {
        res = append(res,tmp[len-i-1])
    }
    return string(res)
}