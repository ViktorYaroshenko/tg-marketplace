package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewStockAddEvent(t *testing.T) {
	requestedAt := time.Unix(1795801168, 0).UTC()

	event := newStockAddEvent(
		12345, 100, "Яблоки", 10, requestedAt,
	)

	if _, err := uuid.Parse(event.EventID); err != nil {
		t.Fatalf("invalid event ID: %v", err)
	}

	if event.ProductName != "Яблоки" {
		t.Error("unexpected product name")
	}

	if event.Quantity != 10 {
		t.Error("unexpected quantity")
	}

	if !event.RequestedAt.Equal(requestedAt) {
		t.Error("unexpected timestamp")
	}
}

func TestEventIDDeterministic(t *testing.T) {
	requestedAt := time.Unix(1795801168, 0).UTC()

	first := newStockAddEvent(
		12345, 100, "Яблоки", 10, requestedAt,
	)

	second := newStockAddEvent(
		12345, 100, "Яблоки", 10, requestedAt,
	)

	third := newStockAddEvent(
		12345, 101, "Яблоки", 10, requestedAt,
	)

	if first.EventID != second.EventID {
		t.Error("same update must produce same event ID")
	}

	if first.EventID == third.EventID {
		t.Error("different updates must produce different event IDs")
	}
}
