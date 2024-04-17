package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQEventProvider struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	exchange   string
}

func MakeRabbitMQEventProvider(connectionURL string, exchangeName string) *RabbitMQEventProvider {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic("Could not connect to RabbitMQ")
		return nil
	}

	channel, err := conn.Channel()

	if err != nil {
		panic("Could not create channel")
		return nil
	}

	err = channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)

	if err != nil {
		panic("Could not declare exchange")
		return nil
	}

	return &RabbitMQEventProvider{
		connection: conn,
		channel:    channel,
		exchange:   exchangeName,
	}

}

func (r *RabbitMQEventProvider) ConsumeEvents() {

	queue, channelError := r.channel.QueueDeclare("", false, false, false, false, nil)

	if channelError != nil {
		panic("Could not declare queue")

	}

	err := r.channel.QueueBind(queue.Name, "email", "opencred-dev", false, nil)

	if err != nil {
		panic("Could not bind queue")

	}

	msgs, err := r.channel.Consume(queue.Name, "", true, false, false, false, nil)

	if err != nil {
		panic("Could not consume messages")

	}

	log.Println("[x] Waiting for messages [x]")

	forever := make(chan bool)

	go func() {
		for message := range msgs {
			log.Println("[x] Received Message: " + "[x]")

			payloadString := string(message.Body)

			log.Println(payloadString)
		}
	}()
	<-forever

}
