package server

import rabbitmq "github.com/GoPicos-Mailing-Service/internal/infra/providers/rabbitMQ"

func InitializeRabbitMQServer() {
	rabbitMQProvider := rabbitmq.MakeRabbitMQEventProvider("amqp://guest:guest@localhost:5672/", "gopicos-dev")

	rabbitMQProvider.ConsumeEvents()
}
