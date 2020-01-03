package main

import "github.com/go-redis/redis"

var redisClient *redis.Client
// CreateRedisClient redis 连接
func CreateRedisClient() *redis.Client {
	redisClinet := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
		PoolSize: 10,
	})
	return redisClinet
}
// Rules 规则信息.
type Rules struct {
	RuleID int32 `json:"rule_id"`
	RuleType string `json:"rule_type"`
	RuleWord string `json:"keyword"`
	AddTime string `json:"add_time"`
}
// Article 文章信息.
type Article struct {
	ArticleID int32 `json:"article_id"`
	UserID int32  `json:"user_id"`
	Rules Rules  `josn:"rules"`
}