package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	log.Println("Starting consumer")

	/*** INITIALIZE CONSUMER ***/
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal(err)
	}

	//err = consumer.SubscribeTopics([]string{"cars"}, nil)
	if err = consumer.Subscribe("announces", nil); err != nil {
		log.Fatal(err)
	}

	/*** INITIALIZE MONGODB CLIENT ***/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Panic(err)
	}

	client.StartSession()

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	if err = client.Connect(ctx); err != nil {
		log.Panic(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Panic(err)
	}

	collection := client.Database("reezorcar").Collection("announces")

	fmt.Println("HANDLE MESSAGES")
	/*** HANDLE MESSAGES ***/
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			// Handle message on a go routine
			go handleStandalone(string(msg.Value), collection)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
