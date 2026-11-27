package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func newProducer(broker, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:            kafka.TCP(broker),
		Topic:           topic,
		Balancer:        &kafka.Hash{},
		RequiredAcks:    kafka.RequireAll,
		MaxAttempts:     3,
		WriteTimeout:    5 * time.Second,
		ReadTimeout:     5 * time.Second,
		WriteBackoffMin: 100 * time.Millisecond,
		WriteBackoffMax: 1 * time.Second,
		BatchTimeout:    10 * time.Millisecond,
	}
}

func publishEvent(
	ctx context.Context,
	writer *kafka.Writer,
	key string,
	value []byte,
) error {
	return writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: value,
	})
}
