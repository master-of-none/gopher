package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.Subscribe("Test", nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received %s\n", string(msg.Value))
		}
	}
}
