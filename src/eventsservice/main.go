package main

import (
	"flag"
	"github.com/user/rest-api/lib/configuration"
	"fmt"
	"github.com/user/rest-api/lib/persistence/dblayer"
	"github.com/user/rest-api/src/eventsservice/rest"
	"log"
	"github.com/streadway/amqp"
	"os"
)

func main() {
	amqpURL := os.Getenv("AMQP_URL");
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672"
	}
	connection, err := amqp.Dial(amqpURL)
	if err != nil {
		panic("Could not establish AMQP connection: " + err.Error())
	}

	defer connection.Close()
	channel, err := connection.Channel()
	if err != nil {
		panic("Could not open channel: " + err.Error())
	}
	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	message := amqp.Publishing{Body: []byte("Hello World")}
	err = channel.Publish("events", "some-routing-key", false, false, message)
	if err != nil {
		panic("Error while publishing message: " + err.Error())
	}
	_, err = channel.QueueDeclare("my_queue", true, false, false, false, nil)
	if err != nil {
		panic("error while declaring the queue: " + err.Error())
	}
	err = channel.QueueBind("my_queue", "#", "events", false, nil)
	if err != nil {
		panic("error while binding the queue: " + err.Error())
	}
	msgs, err := channel.Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		panic("Error while consuming the queue: " + err.Error())
	}
	for msg := range msgs {
		fmt.Println("Message received: " + string(msg.Body))
		msg.Ack(false)
	}
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to configuration json file")
	flag.Parse()

	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connection to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}

