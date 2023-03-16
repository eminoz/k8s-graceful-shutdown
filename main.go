package main

import (
	"fmt"

	"github.com/eminoz/graceful/broker"
	"github.com/eminoz/graceful/consumer"
)

func main() {
	fmt.Print("started")
	broker.RabbitConnection()
	consumer.RunConsumer()
}
