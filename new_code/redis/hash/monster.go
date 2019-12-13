package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Monster struct {
	Name string
	Age int
	Sex string
}

func main()  {

	var name string
	fmt.Println("输入姓名:")
	fmt.Scanf("%s\n", &name)
	//fmt.Println(name)

	var age int
	fmt.Println("输入年龄:")
	fmt.Scanf("%d", &age)
	//fmt.Println(age)
	var sex string
	fmt.Println("输入性别:")
	fmt.Scanf("%s\n", &sex)

	monster := Monster{
		Name:name,
		Age:age,
		Sex:sex,
	}

	fmt.Println(monster)
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("redis 链接失败:",err)
		return
	}
	defer conn.Close()

	monsterString ,err := json.Marshal(monster)
	if err != nil {
		fmt.Println("monster json.Marshal:",err)
		return
	}

	// 写redis 字符串 string
	_, err = conn.Do("setex","go:monster:string",20,string(monsterString))
	if err != nil {
		fmt.Println("redis setex失败:",err)
		return
	}

	//time.Sleep(time.Second * 2)
	ttl,_ := redis.Int(conn.Do("ttl", "go:monster:string"))

	fmt.Println("go:monster:string 的ttl=",ttl)

	res ,_ := redis.String(conn.Do("get", "go:monster:string"))

	var monster2 Monster

	_ = json.Unmarshal([]byte(res), &monster2)
	fmt.Println(monster2)

	// 写hash
	_, _  = conn.Do(
		"hMset",
		"go:monster:hash", "Name", monster.Name, "Age",monster.Age,"sex",monster.Sex,
		)

	// hetall
	monsterHgetall , _ := redis.Strings(conn.Do("hgetall", "go:monster:hash"))

	fmt.Println(monsterHgetall)
	for index,value := range monsterHgetall {
		fmt.Println("field:",index,"value:",value)
	}

	monsterHgetAge , _ := redis.Int(conn.Do("hget", "go:monster:hash", "Age"))
	fmt.Println("hget Age int 类型:",monsterHgetAge)

	monsterHgetName , _ := redis.String(conn.Do("hget", "go:monster:hash", "Name"))
	fmt.Println("hget Name string 类型:",monsterHgetName)


	// list 操作
	





}