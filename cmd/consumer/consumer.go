package main

import (
	"log"
	"poc-rabbitmq/internal/rabbitmq/consumer"
)

func main() {
	err := consumer.ConsumeMessages()
	if err != nil {
		log.Fatalf("Falha ao consumir mensagens: %s", err)
	}
}
