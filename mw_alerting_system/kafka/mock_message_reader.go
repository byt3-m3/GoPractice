package kafka

import (
	"github.com/segmentio/kafka-go"
)

type ReadMockReturn struct {
	Message kafka.Message
	Err     error
}

type MockMessageReader struct {
	ReadMockReturn func() ReadMockReturn
}

func (m MockMessageReader) Read() (kafka.Message, error) {
	res := m.ReadMockReturn()

	return res.Message, res.Err
}
