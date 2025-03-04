package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func init() {
	// Kafka writer setup
	writer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "events",
		Balancer: &kafka.LeastBytes{},
	}
}

func publishEvent(w http.ResponseWriter, r *http.Request) {
	message := []byte("Hello from Go Kafka Producer!")

	fmt.Println("🟡 Sending message:", string(message)) // Debug log

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
	if err != nil {
		log.Println("❌ Failed to publish:", err)
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		return
	}

	fmt.Println("✅ Event published successfully!") // Debug log
	w.Write([]byte("Event published!"))
}

func main() {
	http.HandleFunc("/publish", publishEvent)

	fmt.Println("🚀 Server started on port 8080") // Debug log
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("❌ Server error:", err)
	}
}
