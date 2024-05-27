package main

import (
	_ "testz/routers"

	beego "github.com/beego/beego/v2/server/web"
	// "context"
	// "fmt"
	// "time"
	// amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	//подключение к серверу RabbitMQ
	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// fmt.Printf(" conn %s\n", err)
	// failOnError(err, "Failed to connect to RabbitMQ")

	// defer conn.Close()
	// //Далее мы создадим канал, в котором находится большая часть API для выполнения задач:
	// ch, err := conn.Channel()
	// fmt.Printf(" ch %s\n", err)
	// failOnError(err, "Failed to open a channel")

	// defer ch.Close()
	// //Для отправки мы должны объявить очередь для отправки; затем мы можем опубликовать сообщение в очереди:
	// q, err := ch.QueueDeclare(
	// 	"hello", // name
	// 	false,   // durable
	// 	false,   // delete when unused
	// 	false,   // exclusive
	// 	false,   // no-wait
	// 	nil,     // arguments
	// )
	// fmt.Printf(" QueueDeclare %s\n", err)
	// failOnError(err, "Failed to declare a queue")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// body := "Hello World!"
	// err = ch.PublishWithContext(ctx,
	// 	"",     // exchange
	// 	q.Name, // routing key
	// 	false,  // mandatory
	// 	false,  // immediate
	// 	amqp.Publishing{
	// 		ContentType: "text/plain",
	// 		Body:        []byte(body),
	// 	})
	// fmt.Printf(" PublishWithContext %s\n", err)
	// failOnError(err, "Failed to publish a message")
	// fmt.Printf(" body %s\n", body)
	// log.Printf(" [x] Sent %s\n", body)

	beego.Run()
}

//Нам также нужна вспомогательная функция, чтобы проверить возвращаемое значение для каждого протокол AMQP
// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Panicf("%s: %s", msg, err)
// 	}
// }
