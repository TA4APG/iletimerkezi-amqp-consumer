package src

import (
	"github.com/streadway/amqp"
)

var Channel *amqp.Channel

func connect() {
	conn, err := amqp.Dial(GetConfigString("AMQP_CONNECTIONSTRING"))
	FailOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	Channel, err = conn.Channel()
	FailOnError(err, "Failed to open a channel")
	// defer Channel.Close()
}
func AmqpConfigure() {
	connect()
	_, err := Channel.QueueDeclare(
		GetConfigString("AMQP_QUEUE_NAME"),
		GetConfigBool("AMQP_QUEUE_DURABLE"),
		GetConfigBool("AMQP_QUEUE_AUTO_DELETE"),
		GetConfigBool("AMQP_QUEUE_EXCLUSIVE"),
		GetConfigBool("AMQP_QUEUE_NO_WAIT"),
		nil)
	FailOnError(err, "Failed to declare an queue")
}
