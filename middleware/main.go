package main

import (
	"fmt"
	"log"
)

type Handler func(event Event) error
type middleware func(next Handler) Handler

var (
	eventTypeCreate eventType = "create"
)

type eventType string
type Event struct {
	payload   interface{}
	eventType eventType
}

type Car struct {
	Make  string
	Model string
}

type EventHandler interface {
	handle(h Handler, middlewares ...middleware) Handler
}

type createCarEvent struct {
	Make  string
	Model string
}

var carBuilderHandler Handler = func(event Event) error {

	switch event.eventType {
	case eventTypeCreate:
		eventPayload := event.payload.(createCarEvent)
		fmt.Println("created", eventPayload)

	}

	return nil
}

func metricMiddleware(next Handler) Handler {

	return func(event Event) error {
		log.Println("incremented metric middleware:", event.eventType)
		return next(event)
	}

}

func logMiddleware(next Handler) Handler {

	return func(event Event) error {
		log.Println("log middleware invoked:", event.eventType)
		return next(event)
	}

}

type eventHandler struct {
}

func (h eventHandler) handle(e Event, handler Handler, middlewares ...middleware) error {
	mws := reverseIterate(middlewares)

	for _, mw := range mws {
		handler = mw(handler)
	}

	if err := handler(e); err != nil {
		return err
	}

	return nil
}

func reverseIterate(middlewares []middleware) []middleware {
	var reveredMiddlewares []middleware
	for i := len(middlewares) - 1; i >= 0; i-- {

		reveredMiddlewares = append(reveredMiddlewares, middlewares[i])
	}

	return reveredMiddlewares
}
func main() {
	eh := eventHandler{}

	e := Event{
		payload: createCarEvent{
			Make:  "bmw",
			Model: "m5",
		},
		eventType: eventTypeCreate,
	}

	if err := eh.handle(e, carBuilderHandler, metricMiddleware, logMiddleware); err != nil {
		log.Fatalln(err)
	}

}
