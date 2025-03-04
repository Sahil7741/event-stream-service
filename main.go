package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

// Struct to parse incoming JSON
type Event struct {
	Message string `json:"message"`
}

func publishEvent(w http.ResponseWriter, r *http.Request) {
	var message string

	// Check if the request has a body (JSON payload)
	if r.Body != nil {
		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Attempt to parse JSON if body is not empty
		var event Event
		if err := json.Unmarshal(body, &event); err == nil {
			// If the JSON is valid and contains a "message" field
			if event.Message != "" {
				message = event.Message
			} else {
				// If JSON is valid but "message" is empty, use default message
				message = "Default message because 'message' was empty in the JSON."
			}
		} else {
			// If JSON parsing fails, fall back to default message
			message = "Hello from Go Kafka Producer!"
		}
	} else {
		// If there's no body, use the default message
		message = "Hello from Go Kafka Producer!"
	}

	// Send message to Kafka
	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(message),
	})
	if err != nil {
		log.Println("❌ Failed to publish:", err)
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		return
	}

	fmt.Println("✅ Event published:", message)
	w.Write([]byte("Event published successfully!"))
}

func main() {
	http.HandleFunc("/publish", publishEvent)

	fmt.Println("🚀 Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("❌ Server error:", err)
	}
}



// Request Samples
// curl -X POST http://localhost:8080/publish

// curl -X POST http://localhost:8080/publish \
//      -H "Content-Type: application/json" \
//      -d '{"message": "Hello, Kafka from API!"}'

// curl -X POST http://localhost:8080/publish \
//   -H "Content-Type: application/json" \
//   -d '{}'

