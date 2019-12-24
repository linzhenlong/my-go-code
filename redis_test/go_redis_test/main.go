package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {

	client := createClient()
	defer client.Close()

	var redisKey string

	keywords := []string{
		"手机",
		"电脑数码",
		"蓝牙耳机",
		"个护化妆",
		"手机通讯",
		"每日白菜",
		"iphone 10",
		"小米数据线",
	}

	redisKey = "user_search_history:" + GenMd5("123456")
	finshed := make(chan int, 2)
	go func() {
		client.Pipeline()
		for _, v := range keywords {
			score := time.Now().UnixNano()
			z := redis.Z{
				Score:  float64(score),
				Member: v,
			}
			fmt.Println(client.ZAdd(redisKey, z))
		}
		finshed<-1
	}()

	 v, ok := <- finshed
	 if ok  {
	 	fmt.Println(v)
		 res := client.ZRevRange(redisKey, 0, -1)
		 fmt.Println(res)
	 }

}


// md5返回字符串
func GenMd5(str string) string  {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return client
}


