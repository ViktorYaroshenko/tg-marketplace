package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type StockAddRequested struct {
	EventID     string    `json:"event_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	RequestedAt time.Time `json:"requested_at"`
}

func newStockAddEvent(
	botID int64,
	updateID int64,
	productName string,
	quantity int,
	requestedAt time.Time,
) StockAddRequested {
	key := fmt.Sprintf(
		"telegram:%d:%d:%d:stock.add.requested",
		botID,
		updateID,
		requestedAt.Unix(),
	)

	eventID := uuid.NewSHA1(uuid.NameSpaceOID, []byte(key))

	return StockAddRequested{
		EventID:     eventID.String(),
		ProductName: productName,
		Quantity:    quantity,
		RequestedAt: requestedAt.UTC(),
	}
}
