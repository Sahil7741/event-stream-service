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
	numWorkers = 30           // Number of workers (consumer goroutines)
	topic      = "events"    // Kafka topic
	broker     = "localhost:9092"
	groupID    = "event-consumer-group" // Single group ID for all workers
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go consumeMessages(i, &wg)
	}

	wg.Wait() // Block indefinitely
}

func consumeMessages(workerID int, wg *sync.WaitGroup) {
	defer wg.Done()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID, // ✅ Only GroupID, no Partition
		MinBytes: 10e3,    // 10KB
		MaxBytes: 10e6,    // 10MB
	})
	defer reader.Close()

	fmt.Printf("👷 Worker %d started...\n", workerID)

	count := 0
	startTime := time.Now()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("❌ Worker %d failed to read: %v", workerID, err)
			return
		}

		fmt.Printf("👷 Worker %d processing: %s\n", workerID, string(msg.Value))
		count++

		// Print throughput every 100 messages
		if count%100 == 0 {
			elapsed := time.Since(startTime).Seconds()
			fmt.Printf("⚡ Worker %d processed %d messages in %.2f sec (%.2f messages/sec)\n",
				workerID, count, elapsed, float64(count)/elapsed)
		}
	}
}
