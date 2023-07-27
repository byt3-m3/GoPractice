package kafka

import "github.com/segmentio/kafka-go"

type MessageReader interface {
	Read() (kafka.Message, error)
}

type messageReader struct {
	conn *kafka.Conn
}

func NewMessageReader(conn *kafka.Conn) MessageReader {
	return &messageReader{
		conn: conn,
	}
}

func (m *messageReader) Read() (kafka.Message, error) {
	return m.conn.ReadMessage(4096)
}
