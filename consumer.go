package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Create a new Kafka reader (subscriber)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "events",
		GroupID:  "event-consumer-group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	fmt.Println("📡 Listening for messages on topic 'events'...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("❌ Failed to read message: %v", err)
		}

		fmt.Printf("📥 Received message: %s\n", string(msg.Value))
	}
}
