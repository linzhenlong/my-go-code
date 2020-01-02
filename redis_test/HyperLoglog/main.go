package main

import (
	"flag"
	"os"
	"time"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

/*
redis HyperLogLog 不精准去重

https://blog.csdn.net/weixin_30394669/article/details/99536073
*/
func CreateRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "",
		PoolSize: 10,
	})
	return redisClient
}

func main() {

	userNum := flag.Int("userNum", 10000, "用户个数")
	//logPath := flag.String("logPath","/tmp/hll.log", "日志")
	flag.Parse()
	log := logrus.Logger{
		Level: logrus.DebugLevel,
		Out:   os.Stdout,
		Formatter: &logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05.000",
		},
	}

	/*fd , err := os.OpenFile(*logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0655)
	if err == nil {
		log.Out = fd
		defer fd.Close()
	}*/

	log.Infof("用户数:%d", *userNum)

	redisClient := CreateRedis()
	f, _ := redisClient.Del("hll:users").Result()
	log.Infof("redisClient.Del", f)
	for i := 1; i <= *userNum; i++ {
		r, err := redisClient.PFAdd("hll:users", i).Result()
		if err != nil {
			log.Error("redisClient.PFAdd error", err, i)
			break
		}

		count, _ := redisClient.PFCount("hll:users").Result()
		if i == *userNum {
			cha := count - int64(i)
			rate := float64(cha / int64(i))
			log.Infof("pfcount=%d,i=%d,r=%d, 差值:%d,概率%f", count, i, r, cha, rate)
		}
	}
	time.Sleep(time.Second * 100)
}
