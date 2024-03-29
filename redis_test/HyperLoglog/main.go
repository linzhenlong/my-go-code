package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"os"
	"strconv"
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

	userNum := flag.Int("userNum", 1000000, "用户个数")
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
	log.Infof("redisClient.Del %d", f)
	/* var (
		sucN int
		errN int
	) */

	for i := 0; i < *userNum; i++ {

		v := strconv.Itoa(i + *userNum)
		md5 := GenMd5(v)
		r, err := redisClient.PFAdd("hll:users", md5).Result()
		if err != nil {
			log.Error("redisClient.PFAdd error", err, i)
			break
		}
		/* if r == 1 {
			sucN++
		} else {
			errN++
		}
		log.Infof("sucN:%d ;errN:%d", sucN, errN) */
		log.Infof("res:%d,i:%s", r, md5)

	}

	time.Sleep(time.Second * 100)
}
func GenMd5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
