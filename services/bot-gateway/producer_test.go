//go:build integration

package main

import (
	"context"
	"testing"
	"time"
)

func TestPublishEvent(t *testing.T) {
	writer := newProducer("localhost:9092", "bot-gateway-test")
	defer writer.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Second,
	)
	defer cancel()

	err := publishEvent(
		ctx,
		writer,
		"test-key",
		[]byte("Hello from bot-gateway"),
	)

	if err != nil {
		t.Fatalf("failed to publish event: %v", err)
	}
}
