package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

// 定义一个全局的pool

var pool *redis.Pool

// 当启动程序时，就初始化连接池
func init()  {
	pool = &redis.Pool{
		MaxIdle:8,// 最大空闲链接数
		MaxActive:0, // 最大连接数，0表示没有限制
		IdleTimeout: time.Second * 100,// 最大空闲时间
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main()  {

	// 从链接池中获取链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "go:pool:key", "链接池111")
	if err != nil {
		panic(err)
	}
	r, err :=redis.String(conn.Do("get", "go:pool:key"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)


	// 如果我们从pool中取出链接，一定要保证连接池是没有关闭的.
	pool.Close()
	conn2 := pool.Get()
	fmt.Println(conn2)
	_, err = conn2.Do("set", "go:pool:key", "链接池111")
	if err != nil {
		panic(err)
	}
	r2, err :=redis.String(conn2.Do("get", "go:pool:key"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r2)


}
