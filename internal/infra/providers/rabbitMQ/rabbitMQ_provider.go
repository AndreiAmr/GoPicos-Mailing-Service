package rabbitmq

import (
	"encoding/json"
	"log"

	triggerMailling "github.com/GoPicos-Mailing-Service/internal/modules/TriggerMailings/controllers"
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

	}

	channel, err := conn.Channel()

	if err != nil {
		panic("Could not create channel")

	}

	err = channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)

	if err != nil {
		panic("Could not declare exchange")

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

	err := r.channel.QueueBind(queue.Name, "mailing", "gopicos-dev", false, nil)

	if err != nil {
		panic("Could not bind queue")

	}

	msgs, err := r.channel.Consume("mailing", "", true, false, false, false, nil)

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

			var payload triggerMailling.TriggerMaillingDTO

			err := json.Unmarshal([]byte(payloadString), &payload)

			if err != nil {
				panic("Could not unmarshal payload")
			}

			triggerMailling.TriggerMailling(triggerMailling.TriggerMaillingDTO{
				Email: payload.Email,
				Token: payload.Token,
			})
		}
	}()
	<-forever

}
