package main

import (
	"log"
	"net/http"

	"github.com/giovannylucas/observability-elastic-stack/golang-server/database"
	"github.com/giovannylucas/observability-elastic-stack/golang-server/routes"
	"github.com/gorilla/handlers"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to the database
	database.ConnectOnDatabase()

	// create a new router
	r := routes.HandleRequests()

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	
	if err != nil {
		log.Println(err)
	}

	log.Println("Connected to RabbitMQ")

	channel, err := connection.Channel()

	messages, err := channel.Consume(
		"command", // queue
		"", 			  // consumer
		true, 		  // auto-ack
		false, 		  // exclusive
		false, 		  // no-local
		false, 		  // no-wait
		nil, 			  // args
	)

	var forever chan struct{}

	go func() {
		for message := range messages {
			log.Printf("Received a message: %s", message.Body)
		}
	}()
	
	defer connection.Close()

	<-forever
	
	// start the server
	log.Println("Server started at port 3334!")
	http.ListenAndServe(":3334", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
}