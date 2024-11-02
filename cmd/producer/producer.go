package main

import (
	"fmt"
	"log"
	"poc-rabbitmq/internal/rabbitmq/publisher"
)

func main() {
	msg := "Hello, RabbitMQ!"
	err := publisher.PublishMessage(msg)
	if err != nil {
		log.Fatalf("Falha ao publicar mensagem: %s", err)
	}
	fmt.Println("Mensagem publicada:", msg)
}
