package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type User struct {
	Name string `json:"name"`
	Age int
	Sex string
	Phone string
}

func main()  {

	var host string = "127.0.0.1:6379"

	var redisHashKeyPrefix string = "go:hash:user"
	conn, err := redis.Dial("tcp", host)

	if err != nil {
		fmt.Printf("redis 链接报错error=%v\n", err)
		return
	}
	defer conn.Close()

	userInfo := make(map[string]string, 5)
	userInfo["name"] = "张三"
	userInfo["age"] = "18"
	userInfo["email"] = "222@sina.com"
	userInfo["phone"] = "13800138000"
	userInfo["sex"] = "男"

	for index, value := range userInfo {
		_, err = conn.Do("hset", redisHashKeyPrefix, index, value)
		if err != nil {
			fmt.Printf("redis hset 写入error=%v\n", err)
			return
		}
	}

	// hgetall
	r , err := redis.Values(conn.Do("hgetall", redisHashKeyPrefix))

	if err != nil {
		fmt.Printf("redis hgetall error=%v\n", err)
		return
	}

	fmt.Println("hgetall:")
	for k, v := range r {
		fmt.Printf("%d-->%s\n",k, v)
	}

	// hget
	r1, err := redis.String(conn.Do("hget",redisHashKeyPrefix, "email"))
	if err != nil {
		fmt.Printf("redis hget error=%v\n", err)
		return
	}
	fmt.Printf("hget email=%s\n",r1)

	// hmget
	r3 , err := redis.Values(conn.Do("hmget", redisHashKeyPrefix,"name","age"))

	if err != nil {
		fmt.Printf("redis hgetall error=%v\n", err)
		return
	}
	fmt.Println("hmget:")
	for k, v := range r3 {
		fmt.Printf("%d-->%s\n",k, v)
	}


	user := User{
		Name:"XXX",
		Age:18,
		Sex:"男",
		Phone:"13800",
	}
	userData ,_ := json.Marshal(user)
	_ ,_= conn.Do("set", "go:struct:user",userData)

	userRes,_ := redis.String(conn.Do("get","go:struct:user"))
	fmt.Println(userRes)

	user2 := User{}
	_ = json.Unmarshal([]byte(userRes),&user2)
	fmt.Println(user2)

}
