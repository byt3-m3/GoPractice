package main

import (
	"log"
)

type BuilderOperation func(car *Car) error

func WithModel(m string) BuilderOperation {
	return func(car *Car) error {
		car.model = m
		log.Println("set model:", m)
		return nil
	}
}

var (
	WithTestData BuilderOperation = func(car *Car) error {
		car.model = "test_model"
		car.make = "test_make"
		car.color = "test_color"
		return nil
	}

	withMake func(make string) BuilderOperation = func(make string) BuilderOperation {
		return func(car *Car) error {
			car.make = make
			return nil
		}
	}
)

func WithColor(color string) BuilderOperation {
	return func(car *Car) error {
		log.Println("set car color:", color)
		car.color = color
		return nil
	}
}

func NewCar(opts ...BuilderOperation) *Car {
	car := &Car{}
	for _, opt := range opts {
		err := opt(car)
		if err != nil {
			log.Fatal(err)
		}
	}

	return car

}

func buildCar(car *Car, operations ...BuilderOperation) *Car {

	for i, _ := range operations {
		opt := operations[i]
		err := opt(car)
		if err != nil {
			log.Fatal(err)
		}
	}

	return car
}

type Car struct {
	color string
	make  string
	model string
}

func (c *Car) GetColor() string {
	return c.color

}

func main() {

	c := NewCar(
		WithTestData,
		WithColor("red"),
		withMake("bmw"),
		WithModel("5-series"),
	)

	log.Println(c)

}
