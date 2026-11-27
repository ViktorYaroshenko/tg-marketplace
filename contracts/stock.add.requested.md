# Stock Add Requested

## Event contract v1

**Status:** Draft

**Topic:** `stock.add.requested`

**Producer:** bot-gateway

**Consumer:** inventory

**Format:** JSON

## Description

A request to increase the stock quantity of a product.

## Example

```json
{
  "event_id": "550e8400-e29b-41d4-a716-446655440000",
  "product_name": "Зелёные яблоки",
  "quantity": 15,
  "requested_at": "2026-11-27T14:33:00Z"
}
```

## Fields

| Field | Type | Required | Description |
|---|---|---|---|
| event_id | UUID | Yes | Unique logical event identifier |
| product_name | string | Yes | Non-empty product name |
| quantity | integer | Yes | Positive quantity |
| requested_at | string | Yes | Original request timestamp in UTC, RFC 3339 |

## Kafka message key

The product name.

## Delivery and idempotency

- Delivery semantics: at-least-once.
- Reprocessing the same Telegram update must preserve the original event ID.
- The consumer must deduplicate events by event_id.
- Deduplication and stock modification must be performed atomically.

## Validation

- product_name must not be empty.
- quantity must be greater than zero.
- event_id must be a valid UUID.
- requested_at must be a valid RFC 3339 timestamp.