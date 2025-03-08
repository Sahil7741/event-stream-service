package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	brokerAddress  = "localhost:9092" // Change to "kafka:9092" if using Docker
	topic          = "events"
	messageCount   = 10000 // Total messages to send
	batchSize      = 1000  // Messages per batch
	concurrency    = 5     // Number of concurrent workers
)

func main() {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		BatchSize: batchSize, // Send messages in batches
	}

	defer writer.Close()

	start := time.Now() // Start time for benchmarking
	var wg sync.WaitGroup
	msgChannel := make(chan kafka.Message, messageCount)

	// Worker function
	sendMessages := func() {
		defer wg.Done()
		for msg := range msgChannel {
			err := writer.WriteMessages(context.Background(), msg)
			if err != nil {
				log.Fatalf("❌ Failed to write message: %v", err)
			}
		}
	}

	// Start concurrent workers
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go sendMessages()
	}

	// Push messages into the channel
	for i := 1; i <= messageCount; i++ {
		msgChannel <- kafka.Message{
			Key:   []byte(fmt.Sprintf("key-%d", i)),
			Value: []byte(fmt.Sprintf("Message %d", i)),
		}

		if i%1000 == 0 {
			fmt.Printf("📤 Sent %d messages...\n", i)
		}
	}

	close(msgChannel) // Close the channel after sending all messages
	wg.Wait()         // Wait for all workers to finish

	elapsed := time.Since(start) // Measure elapsed time
	fmt.Printf("✅ Published %d messages in %.2f seconds (%.2f msgs/sec)\n",
		messageCount, elapsed.Seconds(), float64(messageCount)/elapsed.Seconds())
}
