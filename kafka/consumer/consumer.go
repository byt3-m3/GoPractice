package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/byt3-m3/GoPractice/kafka/common"
	"github.com/segmentio/kafka-go"
	"log"
)

type (
	Config struct {
		KafkaTopic   string
		KafkaBrokers []string
		KafkaReader  *kafka.Reader
	}

	Consumer struct {
		Config Config
	}
)

func (c Consumer) Run(ctx context.Context) error {
	fmt.Println("Starting Consumer")
	for {
		msg, err := c.Config.KafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			log.Fatalln(err)
		}
		fmt.Println("Message Received")

		var event common.Event
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(event)
		payload := unpackPayload(&event)
		fmt.Println(payload)
		if event.Type == "TestEvent" {
			log.Println("Handled Test Event")
		}
	}
	return nil
}

func unpackPayload(event *common.Event) interface{} {
	fmt.Println("Unpacking Payload")
	switch event.Type {
	case "TestEvent":
		var p common.TestEventPayload

		err := json.Unmarshal(event.Payload, &p)
		if err != nil {
			log.Fatalln(err)
		}
		return &p

	}
	return nil
}
