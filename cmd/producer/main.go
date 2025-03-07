package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":              "all",
		"compression.type":  "snappy",
		"retries":           5,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	topic := "orders"

	for i := range 10 {
		key := fmt.Sprintf("user-%d", i%5) // กระจาย Partition ตาม User ID
		value := fmt.Sprintf("Order ID %d", i)

		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:           []byte(key),
			Value:         []byte(value),
		}, nil)

		if err != nil {
			log.Printf("Failed to send message: %v\n", err)
		} else {
			fmt.Printf("✅ Sent: Key=%s, Value=%s\n", key, value)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
