package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool


func initPool(address string, maxIdle,maxActive int,idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:maxIdle, // 8个最大空闲连接
		MaxActive:maxActive, // 最大连接数，0表示没有限制
		IdleTimeout:idleTimeout, // 最大空闲时间
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", address)
		},
	}
}
