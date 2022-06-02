package main

import (
	"context"
	"github.com/byt3-m3/GoPractice/kafka/common"
	"github.com/byt3-m3/GoPractice/kafka/producer"
)

func main() {
	bc := common.BrokerConfig{
		Address: "192.168.1.5",
		Port:    "9092",
	}

	pCfg := producer.NewConfig(bc, "test-topic-3", 0)

	p := producer.Producer{Config: pCfg}
	payload := common.TestEventPayload{
		"This is a test message",
	}

	event := common.BuildEvent("TestEvent", payload)

	p.SendEvent(context.Background(), "some-key", event)

}
