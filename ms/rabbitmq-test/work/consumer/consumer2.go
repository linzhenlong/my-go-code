package main

import(
	rabbitmq "github.com/linzhenlong/my-go-code/ms/rabbitmq-test/RabbitMQ"

)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("linzl-test")

	rabbitmq.ConsumerSimple()
}