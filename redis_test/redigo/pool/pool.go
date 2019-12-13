package pool

import (
	red "github.com/gomodule/redigo/redis"
	"time"
)

type Redis struct {
	Pool *red.Pool
}

func New()*Redis  {
	return &Redis{}
}

func (redis *Redis)Init()*Redis {
	redis.Pool = &red.Pool{
		//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，
		// 而不被清除，随时处于待命状态。
		MaxIdle: 256,

		// 最大的连接数，表示同时最多有N个连接。0表示不限制
		MaxActive: 0,

		// 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭。
		// 如果设置成0，空闲连接将不会被关闭。应该设置一个比redis服务端超时时间更短的时间
		IdleTimeout: time.Duration(120),

		Dial: func() (red.Conn, error) {
			return red.Dial(
				"tcp",
				"127.0.0.1:6379",
				// 读超时
				red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				// 写超时
				red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				// 连接Redis超时时间
				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				red.DialDatabase(0),
			)
		},
	}
	return redis
}