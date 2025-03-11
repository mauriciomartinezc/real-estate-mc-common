package messaging

import "context"

type MessageBroker interface {
	Init() error
	Publish(ctx context.Context, topic string, key string, message []byte) error
	Close() error
}
