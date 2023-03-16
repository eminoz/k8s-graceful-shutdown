package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/eminoz/graceful/model"
)

func User(ctx context.Context, u chan model.User) {

	ch, conn := GetRabbitChannel()
	q, err := ch.QueueDeclare(
		"createdUser", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		fmt.Print(err)
	}

	defer ch.Close()
	defer conn.Close()
	mailqueueConsume, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack you have to rerun code when this field true
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to register consumer", err)
	}

	for d := range mailqueueConsume {

		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Millisecond):
			user := model.User{}
			json.Unmarshal(d.Body, &user)
			u <- user
			d.Ack(false)
		}

	}

}
