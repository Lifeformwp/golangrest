package amqp

import (
	"github.com/streadway/amqp"
	"encoding/json"
)

type amqpEventEmitter struct {
	connection *amqp.Connection
}

func NewAMQPEventEmitter(conn *amqp.Connection) (EventEmitter, error) {
	emitter := &amqpEventEmitter{connection: conn}
	err := emitter.setup()
	if err != nil {
		return nil, nil
	}
	return emitter, nil
}

func (a *amqpEventEmitter) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	return channel.ExchangeDeclare("events", "topic", true, false, false,false, nil)
}

func (a *amqpEventEmitter) Emit(event Event) error {
	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return err
	}
	chaned, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer chaned.Close()
	msg := amqp.Publishing{Headers: amqpTable{"x-event-name": event.EventName()}, Body: jsonDoc, ContentType: "application/json"}
	return chaned.Publish("events", event.EventName(), false, false, msg)
}
