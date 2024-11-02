package rabbitmq

import (
	"fmt"

	"poc-rabbitmq/internal/config"
)

func ConsumeMessages() {
	conn, err := config.ConnectRabbitMQ()
	config.FailOnError(err, "Falha ao conectar no RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	config.FailOnError(err, "Falha ao abrir canal")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"test_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	config.FailOnError(err, "Falha ao declarar a fila")

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	config.FailOnError(err, "Falha ao registrar o consumidor")

	fmt.Println("Aguardando mensagens...")
	for msg := range msgs {
		fmt.Printf("Mensagem recebida: %s\n", msg.Body)
	}
}
