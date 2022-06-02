package main

import (
	"context"
	"fmt"
	"github.com/byt3-m3/GoPractice/kafka/common"
	"github.com/byt3-m3/GoPractice/kafka/consumer"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {

		bc := common.BrokerConfig{
			Address: "192.168.1.5",
			Port:    "9092",
		}
		bcString := []string{fmt.Sprintf("%s:%s", bc.Address, bc.Port)}
		topic := "test-topic-3"

		cCfg := consumer.Config{
			KafkaTopic:   topic,
			KafkaBrokers: bcString,
			KafkaReader: kafka.NewReader(kafka.ReaderConfig{
				Brokers: bcString,
				GroupID: "1",
				Topic:   topic,
			}),
		}

		c := consumer.Consumer{Config: cCfg}
		log.Fatalln(c.Run(context.Background()))

	}()

	wg.Wait()
}
