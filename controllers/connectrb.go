package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRb(body []byte) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	fmt.Printf(" conn %s\n", err)
	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()
	ch, err := conn.Channel()
	fmt.Printf(" ch %s\n", err)
	failOnError(err, "Failed to open a channel")

	defer ch.Close()
	q, err := ch.QueueDeclare(
		"addUser", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	fmt.Printf(" QueueDeclare %s\n", err)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	fmt.Printf(" PublishWithContext %s\n", err)
	failOnError(err, "Failed to publish a message")
	fmt.Printf(" body %s\n", body)
	log.Printf(" [x] Sent %s\n", body)

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
