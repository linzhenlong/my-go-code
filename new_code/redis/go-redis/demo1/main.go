package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"time"
)

func createClient() *redis.Client  {
	client := redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		Password:"",
		DB:0,
	})

	// ping()方法检测redis 是否链接成功
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis 链接失败：",err)
		return nil
	}

	return client
}

func stringOperation(client *redis.Client)()  {

	//set 第三个参数代表过期时间，如果是0代表没有失效时间
	err := client.Set("go-redis:name", "lzl",time.Second * 10).Err()
	if err != nil {
		panic(err)
	}

	// 获取失效时间
	ttl,err := client.TTL("go-redis:name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(ttl))
	fmt.Println(ttl)
	val, err := client.Get("go-redis:name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	//client.GetSet()

}

func main()  {

	client := createClient()
	defer client.Close()
	stringOperation(client)

}
