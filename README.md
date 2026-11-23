# tg-marketplace

Учебный микросервисный бэкенд маркетплейса на Go с интерфейсом в виде Telegram-бота.

## Сервисы
- `ledger`: счета, переводы с двойной записью, идемпотентность
- `inventory`: остатки, резервирование и списание
- `bot-gateway`: Telegram-бот, вызывает API сервисов и показывает события из Kafka

## Стек
Go, PostgreSQL, Redis, Kafka, Docker Compose, Prometheus, Grafana (позже Kubernetes)

## Запуск
```bash
cp deploy/.env.example deploy/.env
docker compose -f deploy/docker-compose.yml up -d
```
Postgres доступен на `localhost:5433`, Kafka на `localhost:9092`.

## Архитектура
Скоро: схема сервисов и событий.