package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Create Kafka writer (publisher)
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "events",
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	// Message to be published
	message := kafka.Message{
		Key:   []byte("event-key"),
		Value: []byte("Hello, Kafka!"),
	}

	// Write message to Kafka
	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	fmt.Println("✅ Event published successfully! \n")
}