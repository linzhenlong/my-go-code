package main

import (
	"fmt"
	"strconv"
	rabbitmq "github.com/linzhenlong/my-go-code/ms/rabbitmq-test/RabbitMQ"
)

func main() {
	// 1.创建mq实例

	rabbitmq := rabbitmq.NewRabbitMQSimple("linzl-test")

	// 2.发送消息
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSimple("hello world"+strconv.Itoa(i))
	}
	fmt.Println("succ")
}
