package common

type (
	Event struct {
		Type    string `json:"type,omitempty"`
		Payload []byte `json:"payload,omitempty"`
	}

	BrokerConfig struct {
		Address string
		Port    string
	}
	TestEventPayload struct {
		Msg string
	}
)
