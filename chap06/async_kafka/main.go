package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/icrowley/fake"
	"os"
	"os/signal"
	"time"
)

type Message struct {
	Who          string
	TimeAsString string
}

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	brokers := []string{"localhost:9092", "localhost:9093", "localhost:9094"}
	asyncProducer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := asyncProducer.Close(); err != nil {
			panic(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	exitProgram := make(chan struct{})

	var nPublished, nErrors int
	go func() {
		for {
			time.Sleep(5 * time.Second)
			body := Message{fake.FirstName(), time.Now().String()}

			payload, _ := json.Marshal(body)
			msg := &sarama.ProducerMessage{
				Topic: "currentTime",
				Key:   sarama.StringEncoder("aKey"),
				Value: sarama.ByteEncoder(payload),
			}

			select {
			case asyncProducer.Input() <- msg:
				nPublished++
				fmt.Println("Published message", nPublished)
			case err := <-asyncProducer.Errors():
				nErrors++
				fmt.Println("Failed to publish message", nErrors, err)
			case <-signals:
				exitProgram <- struct{}{}
			}

			fmt.Printf("Published %d messages, %d errors\n", nPublished, nErrors)
		}
	}()

	<-exitProgram
}
