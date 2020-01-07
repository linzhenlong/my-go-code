package main

import "log"
// bucket 流控算法.

type ConnLimiter struct {
	conCurrentConn int
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		conCurrentConn: cc,
		bucket: make(chan int, cc),
	}
}

func (connLimiter *ConnLimiter) GetConn() bool {
	if len(connLimiter.bucket) >= connLimiter.conCurrentConn {
		log.Printf("达到访问限制了")
		return false
	}

	// 如果没满，往管道里丢个值
	connLimiter.bucket <- 1
	return true
}

func (connLimiter *ConnLimiter) ReleaseConn() {
	c := <- connLimiter.bucket
	log.Printf("New connection coming:%d",c)
}