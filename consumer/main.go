package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func main () {
	log.Println("Starting consumer")

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id": "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}

	//err = consumer.SubscribeTopics([]string{"cars"}, nil)
	if err = consumer.Subscribe("cars", nil); err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			//fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			go handleStandalone(string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

