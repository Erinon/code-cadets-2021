package bootstrap

import "github.com/streadway/amqp"

func RabbitMq() *amqp.Channel {
	conn, err := amqp.Dial("amqp://localhost:5672/")
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return channel
}
