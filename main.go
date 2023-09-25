package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Kaparouita/models/rabbitmq"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, ch, pb, err := rabbitmq.InitializeRabbitMQ()
	if err != nil {
		log.Fatal("Failed to initialize RabbitMQ")
	}
	defer conn.Close()
	defer ch.Close()

	err = rabbitmq.RegisterConsumer(yourMessageHandlerFunc, ch, pb, "name")
	if err != nil {
		log.Fatal("Failed to register consumer")
	}

	//db := repositories.NewDbRepo()
	//srv := core.NewService(db)
	//srv.GetUser(&models.User{})
	//handler := handlers.NewHandler(srv)
	//server := server.NewService(handler)

	//server.Initialize()
	forever := make(chan bool)
	<-forever
}

type Message struct {
	Name string `json:"name"`
}

func yourMessageHandlerFunc(msgs <-chan amqp.Delivery, ch *amqp.Channel, pb *amqp.Channel) error {
	for msg := range msgs {
		func() {
			notification := &Message{}
			defer msg.Ack(false)
			err := json.Unmarshal(msg.Body, notification)
			if err != nil {
				return
			}

			// Process the notification, e.g., print or log it
			fmt.Printf("Received notification: %+v\n", notification.Name)
			notification.Name = "Hello from the other side"
			rabbitmq.PublishJSONMessage(pb, msg.CorrelationId, msg.ReplyTo, "", notification)
		}()
	}
	return nil
}
