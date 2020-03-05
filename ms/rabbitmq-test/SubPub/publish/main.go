package main

import (
	"log"
	"strconv"
	"time"

	rabbitmq "github.com/linzhenlong/my-go-code/ms/rabbitmq-test/RabbitMQ"
)

func main() {

	log.Printf("start......")
	rabbitmq := rabbitmq.NewRabbitMQPubSub("linzl-test2")

	ticker := time.NewTicker(time.Second)

	after := time.After(time.Second * 60)

	for {
		select {
		case t := <-ticker.C:
			rabbitmq.PublishPub("fuck you" + t.Format("2006-01-02 15:04:05") + "," + strconv.Itoa(t.Nanosecond()))
		case <-after:
			ticker.Stop()
			log.Printf("over.....")
			return
		}
	}

}
