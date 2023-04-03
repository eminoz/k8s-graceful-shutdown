package broker

import (
	"log"

	"github.com/streadway/amqp"
)

var Channel *amqp.Channel
var Conn *amqp.Connection

func RabbitConnection() {
	conn, err := amqp.Dial("amqp://" + "myuser" + ":" + "mypassword" + "@" + "10.105.66.16" + ":" + "5672" + "/")
	// conn, err := amqp.Dial("amqp://" + "eminoz" + ":" + "eminoz" + "@" + "localhost" + ":" + "5672" + "/")

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}

	Channel = ch
	Conn = conn
}
func GetRabbitChannel() (*amqp.Channel, *amqp.Connection) {
	return Channel, Conn
}
