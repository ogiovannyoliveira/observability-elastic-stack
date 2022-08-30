package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/giovannylucas/observability-elastic-stack/golang-server/database"
	"github.com/giovannylucas/observability-elastic-stack/golang-server/models"
	"github.com/giovannylucas/observability-elastic-stack/golang-server/routes"
	"github.com/giovannylucas/observability-elastic-stack/golang-server/services/events"
	"github.com/gorilla/handlers"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to the database
	database.Connect()

	// create a new router
	r := routes.HandleRequests()

	// amqp://guest:guest@localhost:5672/  os.Getenv("AMQP_URL")
	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
	
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
	
	defer connection.Close()

	var forever chan struct{}

	go func() {
		for message := range messages {
			log.Printf("Received a message: %s", message.Body)

			var event models.Event
			err := json.Unmarshal(message.Body, &event)

			if err != nil {
				log.Panic(err)
			}

			events.Create(&event)
			
			log.Printf("Event created: %s", event.Name)
		}
		}()
	
	// start the server
	log.Println("Server started at port 3335!")
	http.ListenAndServe(":3335", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
	
	<-forever
}