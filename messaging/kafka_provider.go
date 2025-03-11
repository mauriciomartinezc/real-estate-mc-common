package messaging

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"strings"
	"time"
)

type KafkaProvider struct {
	writer *kafka.Writer
}

func NewKafkaProvider() (MessageBroker, error) {
	brokersStr := os.Getenv("KAFKA_BROKERS")
	if len(brokersStr) == 0 {
		return nil, fmt.Errorf("KAFKA_BROKERS not set")
	}
	brokers := strings.Split(brokersStr, ",")
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	})
	return &KafkaProvider{writer: writer}, nil
}

func (k *KafkaProvider) Init() error {
	if k.writer == nil {
		return fmt.Errorf("el cliente de Kakfa no está inicializado")
	}
	return nil
}

func (k *KafkaProvider) Publish(ctx context.Context, topic string, key string, message []byte) error {
	msg := kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: message,
		Time:  time.Now(),
	}
	return k.writer.WriteMessages(ctx, msg)
}

func (k *KafkaProvider) Close() error {
	return k.writer.Close()
}
