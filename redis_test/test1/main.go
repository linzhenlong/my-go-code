package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
func main()  {
	c, err := redis.Dial("tcp","127.0.0.1:6379")

	if err !=nil {
		fmt.Println("链接失败")
		return
	}
	defer c.Close()

	_,err = c.Do("SET", "lzl_test", "111");

	lzl_test,err := redis.String(c.Do("GET","lzl_test"))
	fmt.Println(lzl_test)

}

