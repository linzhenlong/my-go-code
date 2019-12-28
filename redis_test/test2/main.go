package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {

	redisCli := *redis.NewClient(&redis.Options{
		Addr:    "127.0.0.1:6379",
		Network: "tcp",
	})

	ret ,_:= redisCli.PFAdd("hll", "z").Result()
	fmt.Println(ret)

	r2, _ := redisCli.ZIncrBy("lzl1",10,"tom").Result()
	fmt.Println(r2)
}





