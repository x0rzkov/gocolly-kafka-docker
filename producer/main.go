package main

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func main () {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		log.Panic(err)
	}

	defer p.Close()

	go func () {
		for e := range p.Events()  {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed : %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to %v\n", ev.TopicPartition)
				}

			}
		}
	} ()

	topic := "cars"

	for _, msg := range []string{"This", "is", "a", "teseeet"} {
		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value: []byte(msg),
		}, nil)
		if err != nil {
			log.Printf("Error when producing %s : %s", msg, err.Error())
		}
	}

	p.Flush(15 * 1000)
}
