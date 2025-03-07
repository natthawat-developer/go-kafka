package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "order-consumer-group",
		"auto.offset.reset": "earliest",
		"enable.auto.commit": false, // ใช้ manual commit
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	topic := "orders"
	c.SubscribeTopics([]string{topic}, nil)

	fmt.Println("📥 Consumer started...")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("✅ Received: Key=%s, Value=%s\n", string(msg.Key), string(msg.Value))

			// Commit Offset หลังจากประมวลผลสำเร็จ
			_, err = c.CommitMessage(msg)
			if err != nil {
				fmt.Printf("⚠️ Commit error: %v\n", err)
			}
		} else {
			fmt.Printf("⚠️ Consumer error: %v\n", err)
		}
	}
}
