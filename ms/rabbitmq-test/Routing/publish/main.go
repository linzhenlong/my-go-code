package main

import (
	"fmt"
	rabbitmq "github.com/linzhenlong/my-go-code/ms/rabbitmq-test/RabbitMQ"
	"strconv"
	"time"
)

func main() {

	mqOne := rabbitmq.NewRabbitMQRouting("my-exchange", "key1")
	mqTwo := rabbitmq.NewRabbitMQRouting("my-exchange", "key2")

	for i := 0; i < 10; i++ {
		mqOne.PublishRouting("hello world key1==>" + strconv.Itoa(i))
		mqTwo.PublishRouting("hello world key2==>" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println("success" + strconv.Itoa(i))
	}
}
