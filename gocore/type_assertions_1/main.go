package main

import "fmt"

type (
	EventId string

	Event struct {
		EventId `json:"eventId,omitempty"`
	}

	PingRequestEvent struct {
		EventId `json:"eventId,omitempty"`
		Target  string `json:"target,omitempty"`
	}
	PingCompleteEvent struct {
		EventId `json:"eventId,omitempty"`
		Data    string `json:"data,omitempty"`
		Target  string `json:"target,omitempty"`
	}
)

func main() {
	event := Event{EventId: "1"}
	pingCompleteEvent := PingCompleteEvent{EventId: "1", Target: "127.0.0.1", Data: "ResultStr"}

	pingRequestEvent := PingRequestEvent{EventId: "1", Target: "127.0.0.1"}

	// Handles event with no panic during assertion
	processEventNoPanic(pingRequestEvent)
	processEventNoPanic(pingCompleteEvent)
	processEventNoPanic(event)

	//// Handles event with  panic during assertion
	//processEventPanic(pingRequestEvent)
	//processEventPanic(pingCompleteEvent)
	//processEventPanic(event)

}

func processEventNoPanic(event interface{}) {
	fmt.Println("Event: ", event)
	completeData, ok := event.(PingCompleteEvent)
	if !ok {
		fmt.Println("Not PingCompleteEvent")
	}
	if completeData.Data == "ResultStr" {
		fmt.Println("Condition Met")
	}
}

func processEventPanic(event interface{}) {
	fmt.Println("Event: ", event)
	completeData := event.(PingCompleteEvent)
	if completeData.Data == "ResultStr" {
		fmt.Println("Condition Met")
	}
}
