package common

import (
	"encoding/json"
	"log"
)

func SerializeEvent(event interface{}) []byte {
	data, err := json.Marshal(event)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func BuildEvent(eventType string, payload interface{}) *Event {

	payloadBytes := SerializeEvent(payload)
	return &Event{
		Type:    eventType,
		Payload: payloadBytes,
	}
}
