package rabbitmq

import (
	"fmt"

	"poc-rabbitmq/internal/config"

	"github.com/streadway/amqp"
)

func PublishMessage(msg string) {
	conn, err := config.ConnectRabbitMQ()
	config.FailOnError(err, "Falha ao conectar no RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	config.FailOnError(err, "Falha ao abrir canal")
	defer ch.Close()

	queue, err := ch.QueueDeclare("test_queue", true, false, false, false, nil)
	config.FailOnError(err, "Falha ao declarar fila")

	err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})

	config.FailOnError(err, "Falha ao publicar mensagem")

	fmt.Println("Mensagem publicada com sucesso!")
}
