package producer

import (
	"context"
	"fmt"
	"github.com/byt3-m3/GoPractice/kafka/common"
	"github.com/segmentio/kafka-go"
	"log"
)

type (
	config struct {
		KafkaConn *kafka.Conn
	}

	Producer struct {
		Config *config
	}
)

func NewConfig(brokerConfig common.BrokerConfig, topic string, partition int) *config {
	conn, err := kafka.DialLeader(context.Background(), "tcp", fmt.Sprintf("%s:%s", brokerConfig.Address, brokerConfig.Port), topic, partition)
	if err != nil {
		log.Fatalln(conn)
	}
	return &config{KafkaConn: conn}
}

func (p Producer) SendEvent(ctx context.Context, key string, event *common.Event) {
	fmt.Println("Sending Event")
	fmt.Println(event)
	valueBytes := common.SerializeEvent(event)

	count, err := p.Config.KafkaConn.WriteMessages(kafka.Message{Key: []byte(key), Value: valueBytes})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Wrote %d lines to broker", count)
}
