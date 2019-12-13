package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main()  {

	var (
		redisKey string
		host string
	)
	redisKey = "go:test:set:key:1"
	host = "127.0.0.1:6379"

	conn, err := redis.Dial("tcp", host)
	if err != nil {
		fmt.Printf("redis 链接报错error=%v\n", err)
		return
	}

	// 关闭连接.
	defer conn.Close()

	// 写入
	_, err = conn.Do("Set", redisKey, "lzl")
	if err !=nil {
		fmt.Printf("go redis set 出错error=%v\n",err)
		return
	}

	// conn.Do("Get",redisKey) 返回r(接口),err,
	// 因为写入的是字符串，因此需要redis.String()转换

	r,err := redis.String(conn.Do("Get",redisKey))
	if err !=nil {
		fmt.Printf("go redis Get 出错error=%v\n",err)
	}
	fmt.Printf("%v\n",r);


	// mset
	_,_ =conn.Do("MSET","go:test:set:key:name","xxxx","go:test:set:key:age",19)

	// mget
	r2,_ := redis.Strings(conn.Do("MGET","go:test:set:key:name","go:test:set:key:age"))
	fmt.Println(r2)
	for _,v := range r2 {
		fmt.Println(v)
	}


}
