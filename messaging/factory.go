package messaging

import "errors"

func NewMessageBroker(nameBroker string) (MessageBroker, error) {
	if nameBroker == "kafka" {
		return NewKafkaProvider()
	}
	return nil, errors.New("message provider not supported")
}
