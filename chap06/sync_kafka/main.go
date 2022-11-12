package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/icrowley/fake"
)

func main() {

	// Config
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Errors = true    // For sync producer this needs to be true
	config.Producer.Return.Successes = true // For sync producer this needs to be true

	// Connect to a Kafka broker running locally
	brokers := []string{"localhost:9092", "localhost:9093", "localhost:9094"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	// cleanup
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: "sync_kafka",
		Value: sarama.StringEncoder(fake.FirstName()),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("FAILED to publish message: %s\n", err)
	} else {
		fmt.Printf("message sent | partition(%d)/offset(%d)\n", partition, offset)
	}
}
