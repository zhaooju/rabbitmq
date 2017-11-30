package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"fmt"
)


const (
	mqurl ="amqp://test1:test1@172.16.10.212:5672"
)


var conn *amqp.Connection
var channel *amqp.Channel

//异常处理
func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

//连接mq
func mqConnect() {
	var err error
	conn, err = amqp.Dial(mqurl)
	failOnErr(err, "failed to connect tp rabbitmq")

	channel, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
}

//发送 返回发送状态
func Push(exchange string,queueName string,message string) bool{

	if channel == nil {
		mqConnect()
	}
	e := channel.Publish(exchange, queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if e != nil {
		return true
	}else{
		return false
	}
}


//监听 返回<-chan amqp.Delivery
func Receive(queueName string) (<-chan amqp.Delivery, error) {
	if channel == nil {
		mqConnect()
	}

	msgs, err := channel.Consume(queueName, "", true, false, false, false, nil)

	failOnErr(err, "")

	return msgs,err

}