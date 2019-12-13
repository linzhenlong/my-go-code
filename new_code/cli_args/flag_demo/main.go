package main

import (
	"flag"
	"fmt"
)

func main()  {
	// 先定义几个变量

	var (
		user string
		pwd string
		port int
		host string
	)
	flag.StringVar(&user,"u", "" ,"获取用户名")
	flag.StringVar(&pwd, "p","", "密码")
	flag.StringVar(&host, "h","", "主机名")
	flag.IntVar(&port, "P", 3306, "端口")


	// 必须调用这个方法进行转换,非常重要，不调用不行.
	flag.Parse()
	fmt.Println("user=",user)
	fmt.Println("pwd=",pwd)
	fmt.Println("port=",port)
	fmt.Println("host=",host)
}
