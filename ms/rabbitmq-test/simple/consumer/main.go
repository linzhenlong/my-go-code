package main

import (
	rabbitmq "github.com/linzhenlong/my-go-code/ms/rabbitmq-test/RabbitMQ"
)

func main() {

	// rabbitmq 实例
	rabbitmq := rabbitmq.NewRabbitMQSimple("linzl-test")

	// 消费
	rabbitmq.ConsumerSimple()
}
