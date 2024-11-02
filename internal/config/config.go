package config

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
