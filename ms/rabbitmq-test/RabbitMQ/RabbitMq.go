package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//MQURL 链接的url. amqp://{用户名}:{密码}@{地址}:{端口号}/{vhost}
const MQURL = "amqp://linzltest:linzltest@127.0.0.1:5672/linzltest"

// RabbitMQ 结构体.
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 链接
	Mqurl string
}

// NewRabbitMq 创建一个rabbitmq实例
func NewRabbitMq(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建链接失败")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}

// Destory 资源断开链接,断开channel和conn.
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理.
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err.Error())
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// NewRabbitMQSimple  简单模式下创建rabbitmq实例
// 简单模式 exchange 及key 都用默认的.
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// exchange 为空，会使用rabbitmq 的default exchange.
	rabbitmq := NewRabbitMq(queueName, "", "")
	return rabbitmq
}

// PublishSimple 简单模式下:队列生产.
func (r *RabbitMQ) PublishSimple(message string) {
	// 1.申请队列，如果队列不存在会自动创建,如果存在则跳过创建
	// 保证队列存在，消费能发送到队列中

	_, err := r.channel.QueueDeclare(
		r.QueueName, // 队列名称
		false,       // 消息是否持久化
		false,       // 是否自动删除
		false,       //是否具有排他性
		false,       // 是否阻塞
		nil,         // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}
	// 2.发送消息到队列中.
	err = r.channel.Publish(
		r.Exchange, //
		r.QueueName,
		false, // 如果是true,会根据exchange类型和routkey规则，如果无法找到符合条件的队列，那么会把发送的消息返还给消费着
		false, // 如果是true,当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		fmt.Println(err)
	}
}

// ConsumerSimple 简单模式下消费队列.
func (r *RabbitMQ) ConsumerSimple() {
	// 1.申请队列，如果队列不存在会自动创建,如果存在则跳过创建
	// 保证队列存在，消费能发送到队列中

	_, err := r.channel.QueueDeclare(
		r.QueueName, // 队列名称
		false,       // 消息是否持久化
		false,       // 是否自动删除
		false,       //是否具有排他性
		false,       // 是否阻塞
		nil,         // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	// 2.接收消息.
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    // 用于区分多个消费者
		true,  // 是否自动应当，默认为true
		false, //是否具有排他性
		false, // 如果设置为true,表示不能讲同一个connection 中发送的消息,传递给同个connection中的消费者
		false, // 是否阻塞
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 消费
	forever := make(chan bool)

	// 3.起携程处理消息
	go func() {
		for d := range msgs {
			// 业务逻辑
			log.Printf("Received a message:%s", d.Body)
			//fmt.Println(d.Body)
		}
	}()
	log.Printf("[*] waiting for message,to exit, press CTRL+C")
	<-forever
}

// NewRabbitMQPubSub 订阅发布模式下的rabbitmq 实例.
func NewRabbitMQPubSub(exchange string) *RabbitMQ {
	rabbitmq := NewRabbitMq("", exchange, "")
	return rabbitmq
}

// PublishPub 订阅模式下的生产
func (r *RabbitMQ) PublishPub(message string) {
	// 1.尝试创建交换机.
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", // 广播类型
		true,
		false,
		false, // 如果是true 表示这个exchange 不可以被client 用了推送消息，仅用来进行exchange 与exchange之间的绑定
		false,
		nil,
	)
	r.failOnErr(err, "failed to declare exchange")

	// 2.发布消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.failOnErr(err, "failed to Publish message")
}

// ConsumerPub 订阅模式下的消费端代码.
func (r *RabbitMQ) ConsumerPub() {
	//1 .试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "failed to declare exchange")

	// 2.试探性创建队列，注意这里的队列名称不需要写
	q, err := r.channel.QueueDeclare(
		"", // 随机队列名称
		false,
		false,
		true,  // 排他性
		false, // 是否阻塞
		nil,
	)
	r.failOnErr(err, "failed to declare queue")

	// 3.绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)
	r.failOnErr(err, "failed to bind queue")

	// 4.消费消息
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			// 业务逻辑
			log.Printf("Received a message:%s", d.Body)
			//fmt.Println(d.Body)
		}
	}()
	log.Printf("[*] waiting for message,to exit, press CTRL+C")
	<-forever
}

// NewRabbitMQRouting 创建路由模式下的rabbitmq 实例.
func NewRabbitMQRouting(exchangeName, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMq("", exchangeName, routingKey)
	return rabbitmq
}

// PublishRouting 路由模式下生产消息.
func (r *RabbitMQ) PublishRouting(message string) {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct", //指定kind 是direct
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "failed to declare exchange")

	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.failOnErr(err, "failed to Publish message")
}

// ConsumerRouting 路由模式下，消息的消费.
func (r *RabbitMQ) ConsumerRouting() {
	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "failed to declare exchange")

	// 2.试探性创建队列
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "failed to declare queue")

	//3 .绑定队列到交换机
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	r.failOnErr(err, "failed to bind queue")

	// 4.读取消息
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			// 业务逻辑
			log.Printf("Received a message:%s", d.Body)
			//fmt.Println(d.Body)
		}
	}()
	log.Printf("[*] waiting for message,to exit, press CTRL+C")
	<-forever
}
